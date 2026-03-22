import { useTranslation } from 'next-i18next';
import Router, { useRouter } from 'next/router';
import { serverSideTranslations } from 'next-i18next/serverSideTranslations';
// utils
import usePrice from '@/utils/use-price';
import { useIsRTL } from '@/utils/locals';
// config
import { Config } from '@/config';
import { Routes } from '@/config/routes';
// hooks
import { useVisitQuery } from '@/data/visit';
import { useCreateBillingMutation } from '@/data/billing';
// types
import { LabTest, PrescribedDrug } from '@/types';
// components
import Button from '@/components/ui/button';
import Card from '@/components/common/card';
import { Table } from '@/components/ui/table';
import Layout from '@/components/layouts/admin';
import Loader from '@/components/ui/loader/loader';
import ErrorMessage from '@/components/ui/error-message';
import { NoDataFound } from '@/components/icons/no-data-found';
import { DownloadIcon } from '@/components/icons/download-icon';
import BillingViewHeader from '@/components/billing/billing-view-header';

export default function UpdateVisitPage() {
  const { t } = useTranslation();
  const router = useRouter();
  const { query, locale } = useRouter();
  const { alignLeft, alignRight } = useIsRTL();

  // query
  const {
    visit,
    isLoading: loading,
    error,
  } = useVisitQuery({
    slug: query.id as string,
    language:
      query.action!.toString() === 'edit' ? locale! : Config.defaultLanguage,
  });
  // mutation
  const { mutate: createBill, isLoading } = useCreateBillingMutation();

  if (loading) return <Loader text={t('common:text-loading')} />;
  if (error) return <ErrorMessage message={error.message} />;

   function handleGenerateBill() {
    if (visit) {
 createBill({
        visit_id: visit?.id,
      });

      // Router.push(`${Routes.billing.list}/${response.id}/edit`);
    }
  }

  const drugsTableColumns = [
    {
      title: t('table:table-item-products'),
      dataIndex: 'drug_name',
      key: 'drug_name',
      align: alignLeft,
      render: (drug_name: string, item: PrescribedDrug) => (
        <div>
          <span>{drug_name}</span>
          <span className="mx-2">x</span>
          {/* <span className="font-semibold text-heading">
              {item.pivot.order_quantity}
            </span> */}
        </div>
      ),
    },
    {
      title: t('table:table-item-total'),
      dataIndex: 'price',
      key: 'price',
      align: alignRight,
      render: function Render(_: any, item: PrescribedDrug) {
        const { price } = usePrice({
          amount: item.price,
        });
        return <span>{price}</span>;
      },
    },
  ];

  const labTestsTableColumns = [
    {
      title: t('table:table-item-lab-tests'),
      dataIndex: 'test_name',
      key: 'test_name',
      align: alignLeft,
      render: (test_name: string, item: LabTest) => (
        <div>
          <span>{test_name}</span>
        </div>
      ),
    },
    {
      title: t('table:table-item-total'),
      dataIndex: 'price',
      key: 'price',
      align: alignRight,
      render: function Render(_: any, item: PrescribedDrug) {
        const { price } = usePrice({
          amount: item.price,
        });
        return <span>{price}</span>;
      },
    },
  ];

  return (
    <>
      <Card className="relative overflow-hidden">
        <div className="mb-6 -mt-5 -ml-5 -mr-5 md:-mr-8 md:-ml-8 md:-mt-8">
          <BillingViewHeader
            visitId={visit?.id}
            patientName={visit?.patient?.name}
            wrapperClassName="px-8 py-4"
          />
        </div>
        <div className="flex w-full">
          <Button
            onClick={() => handleGenerateBill()}
            className="mb-5 bg-blue-500 ltr:ml-auto rtl:mr-auto"
            loading={isLoading}
          >
            <DownloadIcon className="h-4 w-4 me-3" />
            {t('common:text-generate')} {t('common:text-bill')}
          </Button>
        </div>

        <div className="mb-10">
          {visit?.prescribed_drugs && (
            <>
              <h2 className="mt-12 mb-5 text-xl font-bold text-heading">
                Recommended Drugs
              </h2>
              <Table
                //@ts-ignore
                columns={drugsTableColumns}
                emptyText={() => (
                  <div className="flex flex-col items-center py-7">
                    <NoDataFound className="w-52" />
                    <div className="mb-1 pt-6 text-base font-semibold text-heading">
                      {t('table:empty-table-data')}
                    </div>
                    <p className="text-[13px]">
                      {t('table:empty-table-sorry-text')}
                    </p>
                  </div>
                )}
                data={visit?.prescribed_drugs}
                rowKey="id"
                scroll={{ x: 300 }}
              />
            </>
          )}
          {visit?.lab_tests && (
            <>
              <h2 className="mt-12 mb-5 text-xl font-bold text-heading">
                Recommended Lab Tests
              </h2>
              <Table
                //@ts-ignore
                columns={labTestsTableColumns}
                emptyText={() => (
                  <div className="flex flex-col items-center py-7">
                    <NoDataFound className="w-52" />
                    <div className="mb-1 pt-6 text-base font-semibold text-heading">
                      {t('table:empty-table-data')}
                    </div>
                    <p className="text-[13px]">
                      {t('table:empty-table-sorry-text')}
                    </p>
                  </div>
                )}
                data={visit?.lab_tests}
                rowKey="id"
                scroll={{ x: 300 }}
              />
            </>
          )}
        </div>
        <div>
          <h2 className="mt-12 mb-5 text-xl font-bold text-heading">Notes</h2>
          <div className="mb-12 flex items-start rounded border border-gray-700 bg-gray-100 p-4">
            {visit?.clinical_notes
              ? visit?.clinical_notes[0]?.note
              : 'No clinical notes found.'}
          </div>
        </div>
      </Card>
    </>
  );
}

UpdateVisitPage.Layout = Layout;

export const getServerSideProps = async ({ locale }: any) => ({
  props: {
    ...(await serverSideTranslations(locale, ['form', 'table', 'common'])),
  },
});
