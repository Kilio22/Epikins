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
import MyProjects from '../components/myProjects/MyProjects';
import BuildLog from '../components/buildLog/BuildLog';

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
        path: routePrefix + 'projects/:module/:project',
        name: '',
        component: ProjectJobs,
        role: 'projects',
        inNavbar: false
    },
    {
        path: routePrefix + 'manage/projects',
        name: 'Projects management',
        component: ProjectsManagement,
        role: 'module',
        inNavbar: true
    },
    {
        path: routePrefix + 'manage/users',
        name: 'Users',
        component: Users,
        role: 'users',
        inNavbar: true
    },
    {
        path: routePrefix + 'manage/credentials',
        name: 'Jenkins credentials',
        component: JenkinsCredentials,
        role: 'credentials',
        inNavbar: true
    },
    {
        path: routePrefix + 'build-log',
        name: 'Build log',
        component: BuildLog,
        role: 'log',
        inNavbar: true
    },
    {
        path: routePrefix + 'student/projects',
        name: 'My projects',
        component: MyProjects,
        role: 'student',
        inNavbar: true
    }
];
