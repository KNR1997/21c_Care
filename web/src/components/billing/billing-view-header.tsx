import cn from 'classnames';
import { useTranslation } from 'next-i18next';
// types
import { Billing, OrderStatus, PaymentGateway, PaymentStatus } from '@/types';
// components
import Badge from '@/components/ui/badge/badge';
import StatusColor from '@/components/order/status-color';

interface BillingViewHeaderProps {
  visitId?: number;
  patientName?: string;
  // order: any;
  // billing?: Billing;
  wrapperClassName?: string;
  buttonSize?: 'big' | 'medium' | 'small';
}

export default function BillingViewHeader({
  visitId,
  patientName,
  wrapperClassName = 'px-11 py-5',
  buttonSize = 'medium',
}: BillingViewHeaderProps) {
  const { t } = useTranslation('common');
  // const isPaymentCOD = [PaymentGateway.COD, PaymentGateway.CASH].includes(
  //   order?.payment_gateway,
  // );
  // const isOrderPending = ![OrderStatus.CANCELLED, OrderStatus.FAILED].includes(
  //   order?.order_status,
  // );
  // const isPaymentActionPending =
  //   !isPaymentCOD &&
  //   isOrderPending &&
  //   order?.payment_status !== PaymentStatus.SUCCESS;

  return (
    <div className={cn(`bg-[#F7F8FA] ${wrapperClassName}`)}>
      <div className="mb-0 flex flex-col flex-wrap items-center justify-between gap-x-8 text-base font-bold text-heading sm:flex-row lg:flex-nowrap">
        <div
          // className={`order-2 flex  w-full gap-6 sm:order-1 max-w-full basis-full justify-between ${
          //   !isPaymentActionPending
          //     ? 'max-w-full basis-full justify-between'
          //     : 'max-w-full basis-full justify-between lg:ltr:mr-auto'
          // }`}
          className='order-2 flex  w-full gap-6 sm:order-1 max-w-full basis-full justify-between'
        >
          <div>
            <span className="mb-2 block lg:mb-0 lg:inline-block lg:ltr:mr-4 lg:rtl:ml-4">
              {t('text-visit-id')} :
            </span>
            <Badge
              text={t(visitId || 'Unknown Visit ID')}
              // color={StatusColor(order?.order_status)}
            />
          </div>
          <div>
            <span className="mb-2 block lg:mb-0 lg:inline-block lg:ltr:mr-4 lg:rtl:ml-4">
              {t('text-patient-name')} :
            </span>
            <Badge
              text={t(patientName || 'Unknown Patient')}
              // color={StatusColor(order?.order_status)}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
