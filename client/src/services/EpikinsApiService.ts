import Axios, { AxiosResponse } from 'axios';
import { apiBaseURI } from '../Config';
import { IWorkgroupsData } from '../interfaces/IWorkgroupsData';
import { IApiUser } from '../interfaces/users/IApiUser';
import { IApiJenkinsCredentials } from '../interfaces/jenkinsCredentials/IApiJenkinsCredentials';
import { IProject } from '../interfaces/projects/IProject';
import { IStudentJob } from '../interfaces/myProjects/IStudentJob';
import { IBuildLogInfo } from '../interfaces/buildLog/IBuildLogInfo';

class EpikinsApiService {
    static async login(accessToken: string): Promise<IApiUser | null> {
        try {
            const res = await Axios.post<IApiUser>(apiBaseURI + '/login', {}, {
                headers: {
                    'Authorization': accessToken
                }
            });
            return res.data;
        } catch (e) {
            console.log(e);
        }
        return null;
    }

    static async getWorkgroupsData(module: string, project: string, city: string, forceUpdate: boolean, apiAccessToken: string): Promise<IWorkgroupsData[] | null> {
        try {
            const res: AxiosResponse<IWorkgroupsData[]> = await Axios.get<IWorkgroupsData[]>(apiBaseURI + '/projects/' + module + '/' + project + '/' + city, {
                headers: {'Authorization': apiAccessToken}, params: {
                    update: forceUpdate
                }
            });
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async getProjects(forceUpdate: boolean, apiAccessToken: string): Promise<IProject[] | null> {
        try {
            const res = await Axios.get<IProject[]>(apiBaseURI + '/projects', {
                headers: {'Authorization': apiAccessToken}, params: {update: forceUpdate}
            });
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async getProjectInformation(project: string, module: string, apiAccessToken: string): Promise<IProject | null> {
        try {
            const res = await Axios.get<IProject>(apiBaseURI + '/projects/' + module + '/' + project, {headers: {'Authorization': apiAccessToken}});
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async getStudentJobs(apiAccessToken: string): Promise<IStudentJob[] | null> {
        try {
            const res = await Axios.get<IStudentJob[]>(apiBaseURI + '/student/jobs', {headers: {'Authorization': apiAccessToken}});
            if (res.data === null) {
                return [];
            }
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async changeProjectBuildLimit(project: IProject, apiAccessToken: string): Promise<IProject[] | null> {
        try {
            const res = await Axios.put<IProject[]>(apiBaseURI + '/projects/' + project.module + '/' + project.job.name, {
                'buildLimit': project.buildLimit
            }, {
                headers: {'Authorization': apiAccessToken}
            });
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async buildStudent(city: string, group: string, project: string, module: string, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.post(apiBaseURI + '/student/build',
                {
                    city,
                    group,
                    module,
                    project: project
                },
                {
                    headers: {'Authorization': apiAccessToken}
                }
            );
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }

    static async buildJobs(requestedBuilds: string[], project: string, visibility: string, city: string, module: string, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.post(apiBaseURI + '/build',
                {
                    city,
                    jobs: requestedBuilds,
                    project,
                    fu: true,
                    module,
                    visibility
                },
                {
                    headers: {'Authorization': apiAccessToken}
                }
            );
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }

    static async globalBuild(project: string, visibility: string, city: string, module: string, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.post(apiBaseURI + '/build/global', {
                    city,
                    module,
                    project,
                    visibility
                },
                {
                    headers: {'Authorization': apiAccessToken}
                }
            );
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }

    static async getUsers(apiAccessToken: string): Promise<IApiUser[] | null> {
        try {
            const res = await Axios.get<IApiUser[]>(apiBaseURI + '/users',
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return res.data;
        } catch (e) {
            console.log(e);
        }
        return null;
    }

    static async updateUser(updatedUser: IApiUser, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.put(apiBaseURI + '/users', updatedUser,
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }

    static async addUser(newUser: IApiUser, apiAccessToken: string): Promise<number> {
        try {
            await Axios.post(apiBaseURI + '/users', newUser,
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return 201;
        } catch (e) {
            return e.response.status;
        }
    }

    static async deleteUser(email: string, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.delete(apiBaseURI + '/users/' + email,
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }

    static async getJenkinsCredentials(apiAccessToken: string): Promise<string[] | null> {
        try {
            const res = await Axios.get<string[]>(apiBaseURI + '/credentials',
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return res.data;
        } catch (e) {
            console.log(e);
        }
        return null;
    }

    static async addJenkinsCredentials(newCredentials: IApiJenkinsCredentials, apiAccessToken: string): Promise<number> {
        try {
            await Axios.post(apiBaseURI + '/credentials', newCredentials,
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return 201;
        } catch (e) {
            return e.response.status;
        }
    }

    static async deleteJenkinsCredentials(login: string, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.delete(apiBaseURI + '/credentials/' + login,
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }

    static async getCities(apiAccessToken: string): Promise<string[] | null> {
        try {
            const res = await Axios.get<string[]>(apiBaseURI + '/cities',
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return res.data;
        } catch (e) {
            console.log(e);
        }
        return null;
    }

    static async getBuildLog(city: string, page: number, projectString: string, starterString: string, apiAccessToken: string): Promise<IBuildLogInfo | null> {
        try {
            const res = await Axios.get<IBuildLogInfo>(apiBaseURI + '/build-log',
                {
                    headers: {'Authorization': apiAccessToken},
                    params: {
                        city,
                        page,
                        project: projectString,
                        starter: starterString
                    }
                });
            return res.data;
        } catch (e) {
            console.log(e);
        }
        return null;
    }

    static async exportBuildLog(city: string, start: number, end: number, project: string, apiAccessToken: string): Promise<any | null> {
        try {
            const res = await Axios.get(apiBaseURI + '/build-log/export',
                {
                    headers: {'Authorization': apiAccessToken},
                    params: {
                        city,
                        start,
                        end,
                        project
                    },
                    responseType: 'blob'
                });
            return res.data;
        } catch (e) {
            console.log(e);
        }
        return null;
    }
}

export default EpikinsApiService;
