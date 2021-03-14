import { ChangeUsersStateByProperty, OnButtonClick } from '../Functions';

export interface IUsersDeletePopupProps {
    onDeleteClick: OnButtonClick,
    changeUsersStateByProperty: ChangeUsersStateByProperty
}

export interface IUsersDeletePopupState {
    isLoading: boolean
}

export const usersDeletePopupInitialState: IUsersDeletePopupState = {
    isLoading: false
};
