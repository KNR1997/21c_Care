import dynamic from 'next/dynamic';
// utils
import { SUPER_ADMIN } from '@/utils/constants';
// components
const AdminLayout = dynamic(() => import('@/components/layouts/admin'));

export default function AppLayout({
  userPermissions,
  ...props
}: {
  userPermissions: string[];
}) {
  if (userPermissions?.includes(SUPER_ADMIN)) {
    return <AdminLayout {...props} />;
  }
}
