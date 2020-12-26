import * as React from 'react';
import { NativeSelect } from '@material-ui/core';
import { IUsersTableBodyProps } from '../../interfaces/users/IUsersTableBody';
import { roles } from './UsersTableHeader';

const UsersTableBody: React.FunctionComponent<IUsersTableBodyProps> = ({
                                                                           modifiedUsers,
                                                                           user,
                                                                           isEditing,
                                                                           jenkinsCredentials,
                                                                           onCheckboxClick,
                                                                           onDeleteClick
                                                                       }) => {
    return (
        <tbody>
        {
            modifiedUsers.map((modifiedUser, modifiedUserIdx) => {
                return (
                    <tr key={modifiedUserIdx}>
                        <td>{modifiedUser.email}</td>
                        {
                            roles.map((role, roleIdx) => {
                                return (
                                    <td key={roleIdx}>
                                        <div className="form-check">
                                            <input type="checkbox"
                                                   className="form-check-input"
                                                   checked={modifiedUser.roles.includes(role.toLocaleLowerCase())}
                                                   disabled={!isEditing}
                                                   onChange={() => onCheckboxClick(modifiedUsers, modifiedUser, modifiedUserIdx, role)}
                                            />
                                        </div>
                                    </td>
                                );
                            })
                        }
                        <td>
                            <NativeSelect variant={'standard'} fullWidth
                                          disabled={!isEditing}
                                          onChange={(event => {
                                              modifiedUsers[modifiedUserIdx].jenkinsLogin = event.target.value;
                                          })}
                                          defaultValue={modifiedUser.jenkinsLogin}>
                                {
                                    jenkinsCredentials.map((value, index) => {
                                        return (
                                            <option key={index}>
                                                {value}
                                            </option>
                                        );
                                    })
                                }
                            </NativeSelect>
                        </td>
                        {
                            <td>
                                <i className={'far fa-trash-alt mr-2 ' + (modifiedUser.email.localeCompare(user.email) === 0 ? 'trash-not-allowed' : 'trash')}
                                   onClick={() => modifiedUser.email.localeCompare(user.email) !== 0 && onDeleteClick(modifiedUser.email)}/>
                            </td>
                        }
                    </tr>
                );
            })
        }
        </tbody>
    );
};

export default UsersTableBody;
