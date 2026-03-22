import { useTranslation } from 'next-i18next';
import { serverSideTranslations } from 'next-i18next/serverSideTranslations';
// components
import Layout from '@/components/layouts/admin';
import CreateOrUpdateLabTestForm from '@/components/labTest/lab-test-form';

export default function CreateLabTestPage() {
  const { t } = useTranslation();
  return (
    <>
      <div className="flex border-b border-dashed border-border-base pb-5 md:pb-7">
        <h1 className="text-lg font-semibold text-heading">
          {t('form:form-title-create-lab-test')}
        </h1>
      </div>
      <CreateOrUpdateLabTestForm />
    </>
  );
}

CreateLabTestPage.Layout = Layout;

export const getStaticProps = async ({ locale }: any) => ({
  props: {
    ...(await serverSideTranslations(locale, ['form', 'common'])),
  },
});
