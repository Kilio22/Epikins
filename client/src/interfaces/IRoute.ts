import { ComponentClass, FunctionComponent } from 'react';
import Projects from '../components/projects/Projects';
import Home from '../components/Home';
import { RouteComponentProps } from 'react-router-dom';
import ProjectJobs from '../components/projectJobs/ProjectJobs';
import Users from '../components/users/Users';
import JenkinsCredentials from '../components/jenkinsCredentials/JenkinsCredentials';
import ProjectsManagement from '../components/projectsManagement/ProjectsManagement';
import { StaticContext } from 'react-router';
import * as H from 'history';

interface IRoute {
    path: string,
    name: string,
    component: ComponentClass<any> | FunctionComponent<any>,
    role: string,
    inNavbar: boolean
}

export interface IRouteProps<PARAMS = {}, S = H.LocationState> {
    routeProps: RouteComponentProps<PARAMS, StaticContext, S>
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
        path: routePrefix + 'manage/users',
        name: 'Users management',
        component: Users,
        role: 'users',
        inNavbar: true
    },
    {
        path: routePrefix + 'manage/credentials',
        name: 'Jenkins credentials management',
        component: JenkinsCredentials,
        role: 'credentials',
        inNavbar: true
    },
    {
        path: routePrefix + 'manage/projects',
        name: 'Projects management',
        component: ProjectsManagement,
        role: 'credentials',
        inNavbar: true
    }
];
