import dynamic from 'next/dynamic';
import { MODAL_VIEWS, useModalAction, useModalState } from './modal.context';
// components
import Modal from '@/components/ui/modal/modal';

const DrugDeleteView = dynamic(
  () => import('@/components/drug/drug-delete-view')
);
const LabTestDeleteView = dynamic(
  () => import('@/components/labTest/lab-test-delete-view')
);
const PatientDeleteView = dynamic(
  () => import('@/components/patient/patient-delete-view'),
);
const SearchModal = dynamic(
  () => import('@/components/layouts/topbar/search-modal'),
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
    case 'SEARCH_VIEW':
      return <SearchModal />;
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
