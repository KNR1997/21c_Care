import { useEffect } from 'react';
import { useTranslation } from 'next-i18next';
import { useLogoutMutation } from '@/data/user';
import { serverSideTranslations } from 'next-i18next/serverSideTranslations';
// components
import Loader from '@/components/ui/loader/loader';

function SignOut() {
  const { t } = useTranslation();
  const { mutate: logout } = useLogoutMutation();

  useEffect(() => {
    logout();
  }, []);

  return <Loader text={t('common:signing-out-text')} />;
}

export default SignOut;

export const getStaticProps = async ({ locale }: any) => ({
  props: {
    ...(await serverSideTranslations(locale, ['common'])),
  },
});
