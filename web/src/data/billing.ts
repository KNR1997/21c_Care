import Router, { useRouter } from 'next/router';
import { useMutation, useQuery, useQueryClient } from 'react-query';
import { toast } from 'react-toastify';
import { useTranslation } from 'next-i18next';
import { Routes } from '@/config/routes';
import { API_ENDPOINTS } from './client/api-endpoints';
import {
  Billing,
  BillingPaginator,
  BillingQueryOptions,
  GetParams,
} from '@/types';
import { mapPaginatorData } from '@/utils/data-mappers';
import { billingClient } from './client/billing';
import { Config } from '@/config';

export const useBillingsQuery = (options: Partial<BillingQueryOptions>) => {
  const { data, error, isLoading } = useQuery<BillingPaginator, Error>(
    [API_ENDPOINTS.BILLINGS, options],
    ({ queryKey, pageParam }) =>
      billingClient.paginated(Object.assign({}, queryKey[1], pageParam)),
    {
      keepPreviousData: true,
    },
  );

  return {
    billings: data?.data ?? [],
    paginatorInfo: mapPaginatorData(data),
    error,
    loading: isLoading,
  };
};

export const useBillingQuery = ({ id }: { id: string }) => {
  const { data, error, isLoading } = useQuery<Billing, Error>(
    [API_ENDPOINTS.CATEGORIES, { id }],
    () => billingClient.get({ slug: id }),
  );

  return {
    billing: data,
    error,
    isLoading,
  };
};

export const useCreateBillingMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(billingClient.create, {
    onSuccess: () => {
      Router.push(Routes.billing.list, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-created'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.BILLINGS);
    },
    onError: (error: any) => {
      toast.error(error?.response?.data?.error)
    }
  });
};

export const useUpdateBillingMutation = () => {
  const { t } = useTranslation();
  const router = useRouter();
  const queryClient = useQueryClient();
  return useMutation(billingClient.update, {
    onSuccess: async (data) => {
      const generateRedirectUrl = router.query.shop
        ? `/${router.query.shop}${Routes.billing.list}`
        : Routes.billing.list;
      await router.push(`${generateRedirectUrl}/${data?.id}/edit`, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-updated'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.BILLINGS);
    },
  });
};

export const useDeleteBillingMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(billingClient.delete, {
    onSuccess: () => {
      toast.success(t('common:successfully-deleted'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.BILLINGS);
    },
  });
};
