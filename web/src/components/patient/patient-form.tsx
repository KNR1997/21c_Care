import { useForm } from 'react-hook-form';
import { useRouter } from 'next/router';
import { Patient } from '@/types';
import { useTranslation } from 'next-i18next';
import { yupResolver } from '@hookform/resolvers/yup';
// validation
import { patientValidationSchema } from './patient-validation-schema';
// hooks
import {
  useCreatePatientMutation,
  useUpdatePatientMutation,
} from '@/data/patient';
// components
import Input from '@/components/ui/input';
import Button from '@/components/ui/button';
import Card from '@/components/common/card';
import Description from '@/components/ui/description';
import SelectInput from '@/components/ui/select-input';
import StickyFooterPanel from '@/components/ui/sticky-footer-panel';

type FormValues = {
  name: string;
  age: string;
  gender: {
    label: string;
    value: string;
  };
};

const defaultValues = {
  name: '',
  age: '',
};

type IProps = {
  initialValues?: Patient | undefined;
};
export default function CreateOrUpdatePatientForm({ initialValues }: IProps) {
  const router = useRouter();
  const { t } = useTranslation();
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
    resolver: yupResolver(patientValidationSchema),
  });

  const { mutate: createPatient, isLoading: creating } =
    useCreatePatientMutation();
  const { mutate: updatePatient, isLoading: updating } =
    useUpdatePatientMutation();

  const onSubmit = async (values: FormValues) => {
    const input = {
      name: values.name,
      age: Number(values.age),
      gender: values.gender.value,
    };
    if (!initialValues) {
      createPatient({
        ...input,
      });
    } else {
      updatePatient({
        ...input,
        id: initialValues.id!,
      });
    }
  };

  const genderOptions = [
    {
      label: 'Male',
      value: 'M',
    },
    {
      label: 'Female',
      value: 'F',
    },
  ];

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className="flex flex-wrap my-5 sm:my-8">
        <Description
          title={t('form:input-label-description')}
          details={`${
            initialValues
              ? t('form:item-description-edit')
              : t('form:item-description-add')
          } ${t('form:patient-description-helper-text')}`}
          className="w-full px-0 pb-5 sm:w-4/12 sm:py-8 sm:pe-4 md:w-1/3 md:pe-5 "
        />

        <Card className="w-full sm:w-8/12 md:w-2/3">
          <Input
            label={t('form:input-label-name')}
            {...register('name')}
            error={t(errors.name?.message!)}
            variant="outline"
            className="mb-5"
            required
          />
          <Input
            label={t('form:input-label-age')}
            {...register('age')}
            error={t(errors.age?.message!)}
            variant="outline"
            className="mb-5"
            required
          />
          {/* <Input
            label={t('form:input-label-gender')}
            {...register('gender')}
            error={t(errors.gender?.message!)}
            variant="outline"
            className="mb-5"
            required
          /> */}
          <SelectInput
            label={t('form:input-label-gender')}
            name="gender"
            control={control}
            // @ts-ignore
            options={genderOptions}
            error={t(errors.gender?.message!)}
            required
          />
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
            disabled={creating || updating}
            className="text-sm md:text-base"
          >
            {initialValues
              ? t('form:button-label-update-patient')
              : t('form:button-label-add-patient')}
          </Button>
        </div>
      </StickyFooterPanel>
    </form>
  );
}
