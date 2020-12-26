import { OnButtonClick } from '../Functions';

export interface IUsersToolboxProps {
    isEditing: boolean,
    isSaving: boolean,
    onSaveClick: OnButtonClick,
    onCancelClick: OnButtonClick,
    onEditClick: OnButtonClick,
    onAddClick: OnButtonClick
}
