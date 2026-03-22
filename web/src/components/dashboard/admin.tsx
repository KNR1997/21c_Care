import { useAnalyticsQuery } from '@/data/dashboard';
import { useTranslation } from 'next-i18next';
import { useEffect, useState } from 'react';

export default function Dashboard() {
  const { t } = useTranslation();
  // const { data, isLoading: loading } = useAnalyticsQuery();
  const [activeTimeFrame, setActiveTimeFrame] = useState(1);
  // const [orderDataRange, setOrderDataRange] = useState(
  //   data?.todayTotalOrderByStatus,
  // );

  // useEffect(() => {
  //   switch (activeTimeFrame) {
  //     case 1:
  //       setOrderDataRange(data?.todayTotalOrderByStatus);
  //       break;
  //     case 7:
  //       setOrderDataRange(data?.weeklyTotalOrderByStatus);
  //       break;
  //     case 30:
  //       setOrderDataRange(data?.monthlyTotalOrderByStatus);
  //       break;
  //     case 365:
  //       setOrderDataRange(data?.yearlyTotalOrderByStatus);
  //       break;

  //     default:
  //       setOrderDataRange(orderDataRange);
  //       break;
  //   }
  // });

  return (
    <div className="grid gap-7 md:gap-8 lg:grid-cols-2 2xl:grid-cols-12">
      <div className="col-span-full rounded-lg bg-light p-6 md:p-7">
        <div className="mb-5 flex items-center justify-between md:mb-7">
          <h3 className="before:content-'' relative mt-1 bg-light text-lg font-semibold text-heading before:absolute before:-top-px before:h-7 before:w-1 before:rounded-tr-md before:rounded-br-md before:bg-accent ltr:before:-left-6 rtl:before:-right-6 md:before:-top-0.5 md:ltr:before:-left-7 md:rtl:before:-right-7 lg:before:h-8">
            {t('text-summary')}
          </h3>
        </div>
      </div>
    </div>
  );
}
