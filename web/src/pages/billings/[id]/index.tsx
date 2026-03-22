import { useRouter } from 'next/router';
import { useTranslation } from 'next-i18next';
import { serverSideTranslations } from 'next-i18next/serverSideTranslations';
// utils
import usePrice from '@/utils/use-price';
import { useIsRTL } from '@/utils/locals';
// hooks
import { useBillingQuery } from '@/data/billing';
import { useDownloadInvoiceMutation } from '@/data/visit';
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

export default function BillingDetailsPage() {
  const { t } = useTranslation();
  const { query } = useRouter();
  const { alignLeft, alignRight } = useIsRTL();
  // query
  const {
    billing,
    isLoading: loading,
    error,
  } = useBillingQuery({
    id: query.id as string,
  });
const { refetch: downloadInvoice } =
  useDownloadInvoiceMutation(billing?.visit_id!);

  const { price: consultation_fee } = usePrice(
    billing && {
      amount: billing?.consultation_fee!,
    },
  );

  const { price: drugs_total } = usePrice(
    billing && {
      amount: billing?.drugs_total!,
    },
  );

  const { price: lab_tests_total } = usePrice(
    billing && {
      amount: billing?.lab_tests_total!,
    },
  );

  const { price: grand_total } = usePrice(
    billing && {
      amount: billing?.grand_total!,
    },
  );

  if (loading) return <Loader text={t('common:text-loading')} />;
  if (error) return <ErrorMessage message={error.message} />;

async function handleDownloadInvoice() {
  const { data } = await downloadInvoice();

  console.log('Invoice data:', data);

  if (data) {
    const url = window.URL.createObjectURL(data);

    const a = document.createElement('a');
    a.href = url;
    a.download = `invoice-${billing?.visit_id}.pdf`;
    document.body.appendChild(a);
    a.click();
    a.remove();

    window.URL.revokeObjectURL(url);
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
            visitId={billing?.visit_id}
            patientName={billing?.patient?.name}
            wrapperClassName="px-8 py-4"
          />
        </div>
        <div className="flex w-full">
          <Button
            onClick={() => handleDownloadInvoice()}
            className="mb-5 bg-blue-500 ltr:ml-auto rtl:mr-auto"
          >
            <DownloadIcon className="h-4 w-4 me-3" />
            {t('common:text-download')} {t('common:text-billing')}
          </Button>
        </div>

        {/* <div className="flex flex-col items-center lg:flex-row">
          <h3 className="mb-8 w-full whitespace-nowrap text-center text-2xl font-semibold text-heading lg:mb-0 lg:w-1/3 lg:text-start">
            {t('form:input-label-visit-id')} - {billing?.visit_id}
          </h3>
        </div> */}

        {/* <div className="my-5 flex items-center justify-center lg:my-10">
          <OrderStatusProgressBox
            orderStatus={'pending' as OrderStatus}
            paymentStatus={'pending' as PaymentStatus}
          />
        </div> */}

        <div className="mb-10">
          {billing?.prescribed_drugs ? (
            <>
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
                data={billing?.prescribed_drugs}
                rowKey="id"
                scroll={{ x: 300 }}
              />
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
                data={billing?.lab_tests}
                rowKey="id"
                scroll={{ x: 300 }}
              />
            </>
          ) : (
            <span>{t('common:no-order-found')}</span>
          )}

          <div className="flex w-full flex-col space-y-2 border-t-4 border-double border-border-200 px-4 py-4 ms-auto sm:w-1/2 md:w-1/3">
            <div className="flex items-center justify-between text-sm text-body">
              <span>{t('common:billing-consultation-fee')}</span>
              <span>{consultation_fee}</span>
            </div>
            <div className="flex items-center justify-between text-sm text-body">
              <span> {t('text-drugs-total')}</span>
              <span>{drugs_total}</span>
            </div>
            <div className="flex items-center justify-between text-sm text-body">
              <span> {t('text-lab-tests-total')}</span>
              <span>{lab_tests_total}</span>
            </div>
            <div className="flex items-center justify-between text-base font-semibold text-heading">
              <span>{t('common:billing-grand-total')}</span>
              <span>{grand_total}</span>
            </div>
          </div>
        </div>
        <div>
          <h2 className="mt-12 mb-5 text-xl font-bold text-heading">Notes</h2>
          <div className="mb-12 flex items-start rounded border border-gray-700 bg-gray-100 p-4">
            {billing?.clinical_notes[0]?.note || 'No clinical notes found.'}
          </div>
        </div>
      </Card>
    </>
  );
}

BillingDetailsPage.Layout = Layout;

export const getServerSideProps = async ({ locale }: any) => ({
  props: {
    ...(await serverSideTranslations(locale, ['form', 'table', 'common'])),
  },
});
