import axios from 'axios';
import Cookies from 'js-cookie';
import { toast } from 'react-toastify';
import { useRouter } from 'next/router';
import { useTranslation } from 'next-i18next';
import { useMutation, useQuery, useQueryClient } from 'react-query';
// utils
import { AUTH_CRED } from '@/utils/constants';
import { setEmailVerified } from '@/utils/auth-utils';
import { mapPaginatorData } from '@/utils/data-mappers';
// config
import { Routes } from '@/config/routes';
// client
import { API_ENDPOINTS } from './client/api-endpoints';
import { userClient } from './client/user';
// types
import { User, QueryOptionsType, UserPaginator } from '@/types';

export const useMeQuery = () => {
  const queryClient = useQueryClient();
  const router = useRouter();

  return useQuery<User, Error>([API_ENDPOINTS.ME], userClient.me, {
    retry: false,

    staleTime: 5 * 60 * 1000, // 5 minutes
    cacheTime: 6 * 60 * 60 * 1000, // 6 hours

    onSuccess: () => {
      if (router.pathname === Routes.verifyEmail) {
        setEmailVerified(true);
        router.replace(Routes.dashboard);
      }
    },

    onError: (err) => {
      if (axios.isAxiosError(err)) {
        if (err.response?.status === 409) {
          setEmailVerified(false);
          router.replace(Routes.verifyEmail);
          return;
        }
        queryClient.clear();
        router.replace(Routes.login);
      }
    },
  });
};

export function useLogin() {
  return useMutation(userClient.login);
}

export const useLogoutMutation = () => {
  const router = useRouter();
  const { t } = useTranslation();

  return useMutation(userClient.logout, {
    onSuccess: () => {
      Cookies.remove(AUTH_CRED);
      router.replace(Routes.login);
      toast.success(t('common:successfully-logout'), {
        toastId: 'logoutSuccess',
      });
    },
  });
};

export const useRegisterMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(userClient.register, {
    onSuccess: () => {
      toast.success(t('common:successfully-register'), {
        toastId: 'successRegister',
      });
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.REGISTER);
    },
  });
};

export const useUpdateUserMutation = () => {
  const { t } = useTranslation();
  const queryClient = useQueryClient();
  return useMutation(userClient.update, {
    onSuccess: () => {
      toast.success(t('common:successfully-updated'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.ME);
      queryClient.invalidateQueries(API_ENDPOINTS.USERS);
    },
  });
};
export const useUpdateUserEmailMutation = () => {
  const { t } = useTranslation();
  const queryClient = useQueryClient();
  return useMutation(userClient.updateEmail, {
    onSuccess: () => {
      toast.success(t('common:successfully-updated'));
    },
    onError: (error) => {
      const {
        response: { data },
      }: any = error ?? {};

      toast.error(data?.message);
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.ME);
      queryClient.invalidateQueries(API_ENDPOINTS.USERS);
    },
  });
};

export const useChangePasswordMutation = () => {
  return useMutation(userClient.changePassword);
};

export const useForgetPasswordMutation = () => {
  return useMutation(userClient.forgetPassword);
};
export const useResendVerificationEmail = () => {
  const { t } = useTranslation('common');
  return useMutation(userClient.resendVerificationEmail, {
    onSuccess: () => {
      toast.success(t('common:PICKBAZAR_MESSAGE.EMAIL_SENT_SUCCESSFUL'));
    },
    onError: () => {
      toast(t('common:PICKBAZAR_MESSAGE.EMAIL_SENT_FAILED'));
    },
  });
};

export const useVerifyForgetPasswordTokenMutation = () => {
  return useMutation(userClient.verifyForgetPasswordToken);
};

export const useResetPasswordMutation = () => {
  return useMutation(userClient.resetPassword);
};

export const useUserQuery = ({ id }: { id: string }) => {
  return useQuery<User, Error>(
    [API_ENDPOINTS.USERS, id],
    () => userClient.fetchUser({ id }),
    {
      enabled: Boolean(id),
    },
  );
};

export const useUsersQuery = (params: Partial<QueryOptionsType>) => {
  const { data, isLoading, error } = useQuery<UserPaginator, Error>(
    [API_ENDPOINTS.USERS, params],
    () => userClient.fetchUsers(params),
    {
      keepPreviousData: true,
    },
  );

  return {
    users: data?.data ?? [],
    paginatorInfo: mapPaginatorData(data as any),
    loading: isLoading,
    error,
  };
};

export const useAdminsQuery = (params: Partial<QueryOptionsType>) => {
  const { data, isLoading, error } = useQuery<UserPaginator, Error>(
    [API_ENDPOINTS.ADMIN_LIST, params],
    () => userClient.fetchAdmins(params),
    {
      keepPreviousData: true,
    },
  );

  return {
    admins: data?.data ?? [],
    paginatorInfo: mapPaginatorData(data as any),
    loading: isLoading,
    error,
  };
};
