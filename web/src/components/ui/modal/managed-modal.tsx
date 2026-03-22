import StoreNoticeDeleteView from '@/components/store-notice/store-notice-delete-view';
import Modal from '@/components/ui/modal/modal';
import dynamic from 'next/dynamic';
import { MODAL_VIEWS, useModalAction, useModalState } from './modal.context';

const DrugDeleteView = dynamic(
  () => import('@/components/drug/drug-delete-view')
);
const LabTestDeleteView = dynamic(
  () => import('@/components/labTest/lab-test-delete-view')
);
const PatientDeleteView = dynamic(
  () => import('@/components/patient/patient-delete-view'),
);
const BanCustomerView = dynamic(
  () => import('@/components/user/user-ban-view'),
);
const UserWalletPointsAddView = dynamic(
  () => import('@/components/user/user-wallet-points-add-view'),
);
const MakeAdminView = dynamic(
  () => import('@/components/user/make-admin-view'),
);
const ApproveShopView = dynamic(
  () => import('@/components/shop/approve-shop-view'),
);
const DisApproveShopView = dynamic(
  () => import('@/components/shop/disapprove-shop-view'),
);
const RemoveStaffView = dynamic(
  () => import('@/components/shop/staff-delete-view'),
);
const ComposerMessage = dynamic(
  () => import('@/components/message/compose-message'),
);
const SearchModal = dynamic(
  () => import('@/components/layouts/topbar/search-modal'),
);
const DescriptionView = dynamic(
  () => import('@/components/shop-single/description-modal'),
);
const TransferShopOwnershipView = dynamic(
  () => import('@/components/shop/transfer-shop-ownership-view'),
);
const OpenAiModal = dynamic(() => import('@/components/openAI/openAI.modal'));

function renderModal(view: MODAL_VIEWS | undefined, data: any) {
  switch (view) {
    case 'GENERATE_DESCRIPTION':
      return <OpenAiModal />;
    case 'DELETE_DRUG':
      return <DrugDeleteView />;
    case 'DELETE_LAB_TEST':
      return <LabTestDeleteView />;
    case 'DELETE_PATIENT':
      return <PatientDeleteView />;
    case 'DELETE_STORE_NOTICE':
      return <StoreNoticeDeleteView />;
    case 'BAN_CUSTOMER':
      return <BanCustomerView />;
    case 'SHOP_APPROVE_VIEW':
      return <ApproveShopView />;
    case 'SHOP_DISAPPROVE_VIEW':
      return <DisApproveShopView />;
    case 'DELETE_STAFF':
      return <RemoveStaffView />;
    case 'MAKE_ADMIN':
      return <MakeAdminView />;
    case 'ADD_WALLET_POINTS':
      return <UserWalletPointsAddView />;
    case 'COMPOSE_MESSAGE':
      return <ComposerMessage />;
    case 'SEARCH_VIEW':
      return <SearchModal />;
    case 'DESCRIPTION_VIEW':
      return <DescriptionView />;
    case 'TRANSFER_SHOP_OWNERSHIP_VIEW':
      return <TransferShopOwnershipView />;
    default:
      return null;
  }
}

const ManagedModal = () => {
  const { isOpen, view, data } = useModalState();
  const { closeModal } = useModalAction();

  return (
    <Modal open={isOpen} onClose={closeModal}>
      {renderModal(view, data)}
    </Modal>
  );
};

export default ManagedModal;
