import * as React from 'react';

export const roles: string[] = [
    'Projects',
    'Users',
    'Credentials',
    'Module'
];

const UsersTableHeader = () => {
    return (
        <thead>
        <tr>
            <th>Email</th>
            {
                roles.map((role, index) => {
                    return <th key={index}>{role}</th>;
                })
            }
            <th>Jenkins login</th>
            <th>Delete</th>
        </tr>
        </thead>
    );
};

export default UsersTableHeader;
