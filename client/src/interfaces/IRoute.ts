import { ComponentClass, FunctionComponent } from 'react';
import Projects from '../components/projects/Projects';
import Home from '../components/Home';
import { RouteComponentProps } from 'react-router-dom';
import ProjectJobs from '../components/projectJobs/ProjectJobs';
import Users from '../components/users/Users';

interface IRoute {
    path: string,
    name: string,
    component: ComponentClass<any> | FunctionComponent<any>,
    role: string,
    inNavbar: boolean
}

export interface IRouteProps<PARAMS> {
    routeProps: RouteComponentProps<PARAMS>
}

export const routePrefix: string = '/';

export const routes: IRoute[] = [
    {
        path: routePrefix,
        name: '',
        component: Home,
        role: '',
        inNavbar: false
    },
    {
        path: routePrefix + 'projects',
        name: 'Projects',
        component: Projects,
        role: 'projects',
        inNavbar: true
    },
    {
        path: routePrefix + 'projects/:project',
        name: '',
        component: ProjectJobs,
        role: 'projects',
        inNavbar: false
    },
    {
        path: routePrefix + 'users',
        name: 'Users',
        component: Users,
        role: 'users',
        inNavbar: true
    }
];
