import { useRouter } from 'next/router';
import { useTranslation } from 'next-i18next';
import { serverSideTranslations } from 'next-i18next/serverSideTranslations';
// config
import { Config } from '@/config';
// hooks
import { useLabTestQuery } from '@/data/labTest';
// components
import Layout from '@/components/layouts/admin';
import Loader from '@/components/ui/loader/loader';
import ErrorMessage from '@/components/ui/error-message';
import CreateOrUpdateLabTestForm from '@/components/labTest/lab-test-form';

export default function UpdateDrugTestCatalogPage() {
  const { t } = useTranslation();
  const { query, locale } = useRouter();
  // query
  const {
    labTest,
    isLoading: loading,
    error,
  } = useLabTestQuery({
    slug: query.id as string,
    language:
      query.action!.toString() === 'edit' ? locale! : Config.defaultLanguage,
  });

  if (loading) return <Loader text={t('common:text-loading')} />;
  if (error) return <ErrorMessage message={error.message} />;

  return (
    <>
      <div className="flex border-b border-dashed border-border-base pb-5 md:pb-7">
        <h1 className="text-lg font-semibold text-heading">
          {t('form:form-title-edit-category')}
        </h1>
      </div>

      <CreateOrUpdateLabTestForm initialValues={labTest} />
    </>
  );
}

UpdateDrugTestCatalogPage.Layout = Layout;

export const getServerSideProps = async ({ locale }: any) => ({
  props: {
    ...(await serverSideTranslations(locale, ['form', 'common'])),
  },
});
