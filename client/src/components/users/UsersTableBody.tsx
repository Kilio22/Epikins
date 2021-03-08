import * as React from 'react';
import { NativeSelect } from '@material-ui/core';
import { IUsersTableBodyProps } from '../../interfaces/users/IUsersTableBody';
import { roles } from './UsersTableHeader';

const UsersTableBody: React.FunctionComponent<IUsersTableBodyProps> = ({
                                                                           users,
                                                                           isEditing,
                                                                           jenkinsCredentials,
                                                                           onCheckboxClick,
                                                                           onFirstDeleteClick
                                                                       }) => {
    const sortedUsers = users.sort(((a, b) => a.email.localeCompare(b.email)));
    return (
        <tbody>
        {
            sortedUsers.map((user, userIdx) => {
                return (
                    <tr key={userIdx}>
                        <td>{user.email}</td>
                        {
                            roles.map((role, roleIdx) => {
                                return (
                                    <td key={roleIdx}>
                                        <div className={'form-check'}>
                                            <input type={'checkbox'}
                                                   className={'form-check-input'}
                                                   checked={user.roles.includes(role.toLocaleLowerCase())}
                                                   disabled={!isEditing}
                                                   onChange={() => onCheckboxClick(users, user, userIdx, role)}
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
                                              users[userIdx].jenkinsLogin = event.target.value;
                                          })}
                                          defaultValue={user.jenkinsLogin}>
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
                                <i className={'far fa-trash-alt mr-2 trash'}
                                   onClick={() => onFirstDeleteClick(user.email)}/>
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
