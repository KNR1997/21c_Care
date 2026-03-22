import { useForm } from 'react-hook-form';
import { useRouter } from 'next/router';
import { LabTest } from '@/types';
import { useTranslation } from 'next-i18next';
import { yupResolver } from '@hookform/resolvers/yup';
// validation
import { labTestValidationSchema } from './lab-test-validation-schema';
// hooks
import {
  useCreateLabTestMutation,
  useUpdateLabTestMutation,
} from '@/data/labTest';
// components
import Input from '@/components/ui/input';
import Button from '@/components/ui/button';
import Card from '@/components/common/card';
import Description from '@/components/ui/description';
import StickyFooterPanel from '@/components/ui/sticky-footer-panel';

type FormValues = {
  name: string;
  default_price: string;
};

const defaultValues = {
  name: '',
  default_price: '',
};

type IProps = {
  initialValues?: LabTest | undefined;
};
export default function CreateOrUpdateLabTestForm({ initialValues }: IProps) {
  const router = useRouter();
  const { t } = useTranslation();
  const {
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
    resolver: yupResolver(labTestValidationSchema),
  });

  const { mutate: createLabTest, isLoading: creating } =
    useCreateLabTestMutation();
  const { mutate: updateLabTest, isLoading: updating } =
    useUpdateLabTestMutation();

  const onSubmit = async (values: FormValues) => {
    const input = {
      name: values.name,
      default_price: Number(values.default_price),
    };
    if (!initialValues) {
      createLabTest({
        ...input,
      });
    } else {
      updateLabTest({
        ...input,
        id: initialValues.id!,
      });
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
          } ${t('form:lab-test-description-helper-text')}`}
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
            label={t('form:input-label-price')}
            {...register('default_price')}
            error={t(errors.default_price?.message!)}
            variant="outline"
            className="mb-5"
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
              ? t('form:button-label-update-lab-test')
              : t('form:button-label-add-lab-test')}
          </Button>
        </div>
      </StickyFooterPanel>
    </form>
  );
}
