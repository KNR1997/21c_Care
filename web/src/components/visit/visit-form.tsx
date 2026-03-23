import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { useRouter } from 'next/router';
import { Patient, Visit } from '@/types';
import { useTranslation } from 'next-i18next';
import { yupResolver } from '@hookform/resolvers/yup';
// validation
import { visitValidationSchema } from './visit-validation-schema';
// hooks
import { usePatientsQuery } from '@/data/patient';
import {
  useCreateVisitMutation,
  usePreviewVisitMutation,
  useUpdateVisitMutation,
} from '@/data/visit';
// components
import Button from '@/components/ui/button';
import Card from '@/components/common/card';
import TextArea from '@/components/ui/text-area';
import Description from '@/components/ui/description';
import SelectInput from '@/components/ui/select-input';
import OpenAIButton from '@/components/openAI/openAI.button';
import ValidationError from '@/components/ui/form-validation-error';
import StickyFooterPanel from '@/components/ui/sticky-footer-panel';
import { toast } from 'react-toastify';

type FormValues = {
  patient: Patient;
  prompt: string;
};

const defaultValues = {
  name: '',
  prompt: '',
};

type IProps = {
  initialValues?: Visit | undefined;
};

export default function CreateOrUpdateVisitForm({ initialValues }: IProps) {
  const router = useRouter();
  const { t } = useTranslation();
  const [aiPreview, setAiPreview] = useState<any>(null);
  const [showPreview, setShowPreview] = useState(false);

  const {
    control,
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>({
    //@ts-ignore
    defaultValues: initialValues
      ? {
          ...initialValues,
        }
      : defaultValues,
    //@ts-ignore
    resolver: yupResolver(visitValidationSchema),
  });

  // mutations
  const { mutateAsync: previewVisit, isLoading: previewing } =
    usePreviewVisitMutation();
  const { mutate: createVisit, isLoading: creating } = useCreateVisitMutation();
  const { mutate: updateVisit, isLoading: updating } = useUpdateVisitMutation();
  // query
  const { patients, loading } = usePatientsQuery({});

  const onSubmit = async (values: FormValues) => {
    const input = {
      patient_id: values.patient.id,
      raw_input: values.prompt,
      ai_result: aiPreview,
    };
    if (!initialValues) {
      createVisit({
        ...input,
      });
    } else {
      updateVisit({
        ...input,
        id: initialValues.id!,
      });
    }
  };

  const handleGeneratePrompt = async () => {
    try {
      const result = await previewVisit({
        raw_input: control._formValues.prompt,
      });
      setAiPreview(result);
      setShowPreview(true);
    } catch (error: any) {
      // Handle error if needed
      toast.error('Failed to generate preview:', error?.response.data.error);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="flex flex-wrap my-5 sm:my-8">
        <Description
          title={t('form:input-label-description')}
          details={`${
            initialValues
              ? t('form:item-description-edit')
              : t('form:item-description-add')
          } ${t('form:visit-description-helper-text')}`}
          className="w-full px-0 pb-5 sm:w-4/12 sm:py-8 sm:pe-4 md:w-1/3 md:pe-5 "
        />

        <Card className="w-full sm:w-8/12 md:w-2/3">
          <div className="mb-5">
            <SelectInput
              label={t('form:input-label-patient')}
              name="patient"
              control={control}
              getOptionLabel={(option: any) => option.name}
              getOptionValue={(option: any) => option.id}
              options={patients!}
              isLoading={loading}
              required
            />
            <ValidationError message={t(errors.patient?.message)} />
          </div>
          <div className="relative mb-5">
            <OpenAIButton
              title={t('form:button-label-analyze-with-ai')}
              onClick={handleGeneratePrompt}
              isLoading={previewing}
            />
            <TextArea
              label={t('form:input-label-prompt')}
              {...register('prompt')}
              error={t(errors.prompt?.message!)}
              variant="outline"
              className="mb-6"
              required
            />
          </div>
          {/* <TextArea
            name="Generated Description"
            inputClassName="h-72"
            label={t('form:input-label-output')}
            // value={data?.result}
            disabled={true}
          /> */}
          {showPreview && aiPreview && (
            <Card className="mt-6">
              <h3 className="text-lg font-semibold mb-4">
                AI Classification Result
              </h3>

              {/* Drugs */}
              <div className="mb-4">
                <h4 className="font-medium">Drugs</h4>
                {aiPreview.drugs?.map((drug: any, index: number) => (
                  <div key={index} className="p-3 border rounded mb-2">
                    {drug.name} - {drug.dosage} - {drug.frequency}
                  </div>
                ))}
              </div>

              {/* Lab Tests */}
              <div className="mb-4">
                <h4 className="font-medium">Lab Tests</h4>
                {aiPreview.lab_tests?.map((test: string, index: number) => (
                  <div key={index} className="p-2 border rounded mb-2">
                    {test}
                  </div>
                ))}
              </div>

              {/* Notes */}
              <div className="mb-4">
                <h4 className="font-medium">Clinical Notes</h4>
                <p>{aiPreview.notes}</p>
              </div>

              {/* <div className="flex justify-end gap-4">
                <Button variant="outline" onClick={() => setShowPreview(false)}>
                  Cancel
                </Button>

                <Button
                  onClick={() =>
                    createVisit({
                      patient_id: control._formValues.patient.id,
                      raw_input: control._formValues.prompt,
                      ...aiPreview,
                    })
                  }
                >
                  Confirm & Save
                </Button>
              </div> */}
            </Card>
          )}
        </Card>
      </div>
      <StickyFooterPanel className="z-0">
        <div className="text-end">
          {initialValues && (
            <Button
              variant="outline"
              onClick={router.back}
              className="text-sm me-4 md:text-base"
              type="button"
            >
              {t('form:button-label-back')}
            </Button>
          )}

          <Button
            loading={creating || updating}
            disabled={creating || updating || !showPreview}
            className="text-sm md:text-base"
          >
            {initialValues
              ? t('form:button-label-update-visit')
              : t('form:button-label-add-visit')}
          </Button>
        </div>
      </StickyFooterPanel>
    </form>
  );
}
