import { useState } from 'react';
import { useTranslation } from 'next-i18next';
// utils
import { useIsRTL } from '@/utils/locals';
// config
import { Routes } from '@/config/routes';
// types
import { MappedPaginatorInfo } from '@/types';
import { Billing, SortOrder, Visit } from '@/types';
// components
import { Table } from '@/components/ui/table';
import Pagination from '@/components/ui/pagination';
import TitleWithSort from '@/components/ui/title-with-sort';
import ActionButtons from '@/components/common/action-buttons';
import { NoDataFound } from '@/components/icons/no-data-found';

export type IProps = {
  billings: Billing[] | undefined;
  paginatorInfo: MappedPaginatorInfo | null;
  onPagination: (key: number) => void;
  onSort: (current: any) => void;
  onOrder: (current: string) => void;
};
const BillingList = ({
  billings,
  paginatorInfo,
  onPagination,
  onSort,
  onOrder,
}: IProps) => {
  const { t } = useTranslation();
  const rowExpandable = (record: any) => record.children?.length;
  const { alignLeft, alignRight } = useIsRTL();
  const [sortingObj, setSortingObj] = useState<{
    sort: SortOrder;
    column: string | null;
  }>({
    sort: SortOrder.Desc,
    column: null,
  });

  const onHeaderClick = (column: string | null) => ({
    onClick: () => {
      onSort((currentSortDirection: SortOrder) =>
        currentSortDirection === SortOrder.Desc
          ? SortOrder.Asc
          : SortOrder.Desc,
      );
      onOrder(column!);

      setSortingObj({
        sort:
          sortingObj.sort === SortOrder.Desc ? SortOrder.Asc : SortOrder.Desc,
        column: column,
      });
    },
  });

  const columns = [
    {
      title: t('table:table-item-visit-id'),
      dataIndex: 'visit_id',
      key: 'visit_id',
      align: alignLeft,
      width: 130,
    },
    {
      title: t('table:table-item-patient'),
      className: 'cursor-pointer',
      dataIndex: 'visit',
      key: 'visit',
      align: alignLeft,
      width: 150,
      render: (visit: Visit) => {
        return (
          <div className="flex flex-wrap gap-1.5 whitespace-nowrap">
            {visit?.patient?.name}
          </div>
        );
      },
    },
    {
      title: (
        <TitleWithSort
          title={t('table:table-item-total')}
          ascending={
            sortingObj.sort === SortOrder.Asc && sortingObj.column === 'slug'
          }
          isActive={sortingObj.column === 'grand_total'}
        />
      ),
      className: 'cursor-pointer',
      dataIndex: 'grand_total',
      key: 'grand_total',
      align: alignLeft,
      width: 150,
    },
    {
      title: t('table:table-item-actions'),
      dataIndex: 'id',
      key: 'actions',
      align: alignRight,
      width: 120,
      render: (id: string, record: Billing) => (
        <ActionButtons
          id={id}
          detailsUrl={Routes.billing.details(id)}
          // deleteModalView="DELETE_BILLING"
        />
      ),
    },
  ];

  return (
    <>
      <div className="mb-6 overflow-hidden rounded shadow">
        <Table
          //@ts-ignore
          columns={columns}
          emptyText={() => (
            <div className="flex flex-col items-center py-7">
              <NoDataFound className="w-52" />
              <div className="mb-1 pt-6 text-base font-semibold text-heading">
                {t('table:empty-table-data')}
              </div>
              <p className="text-[13px]">{t('table:empty-table-sorry-text')}</p>
            </div>
          )}
          data={billings}
          rowKey="id"
          scroll={{ x: 1000 }}
        />
      </div>

      {!!paginatorInfo?.totalRows && (
        <div className="flex items-center justify-end">
          <Pagination
            total={paginatorInfo.totalRows}
            current={paginatorInfo.page}
            pageSize={paginatorInfo.limit}
            onChange={onPagination}
          />
        </div>
      )}
    </>
  );
};

export default BillingList;
