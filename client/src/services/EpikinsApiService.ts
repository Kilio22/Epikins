import Axios, { AxiosResponse } from 'axios';
import { apiBaseURI } from '../Config';
import { IGroupData } from '../interfaces/IGroupData';
import { IJob } from '../interfaces/IJob';
import { IApiUser } from '../interfaces/users/IApiUser';
import { IApiJenkinsCredentials } from '../interfaces/IJenkinsCredentialsTable/IApiJenkinsCredentials';
import { IProject } from '../interfaces/projects/IProject';

class EpikinsApiService {
    static async login(accessToken: string): Promise<IApiUser | null> {
        try {
            const res = await Axios.post<IApiUser>(apiBaseURI + 'login', {}, {
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

    static async getGroupsData(url: string, apiAccessToken: string): Promise<IGroupData[] | null> {
        try {
            const res: AxiosResponse<IGroupData[]> = await Axios.get<IGroupData[]>(url, {headers: {'Authorization': apiAccessToken}});
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async getJobs(url: string, apiAccessToken: string): Promise<IJob[] | null> {
        try {
            const res: AxiosResponse<IJob[]> = await Axios.get<IJob[]>(url, {headers: {'Authorization': apiAccessToken}});
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async getProjects(apiAccessToken: string): Promise<IProject[] | null> {
        try {
            const res = await Axios.get<IProject[]>(apiBaseURI + 'projects', {headers: {'Authorization': apiAccessToken}});
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async buildJobs(requestedBuilds: string[], project: string, visibility: string, fuMode: boolean, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.post(apiBaseURI + 'build',
                requestedBuilds,
                {
                    params: {'visibility': visibility, 'project': project, 'fu': fuMode},
                    headers: {'Authorization': apiAccessToken}
                }
            );
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }

    static async globalBuild(project: string, visibility: string, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.post(apiBaseURI + 'build/global', {},
                {
                    params: {'visibility': visibility, 'project': project},
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
            const res = await Axios.get<IApiUser[]>(apiBaseURI + 'users',
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
            await Axios.put(apiBaseURI + 'users', updatedUser,
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
            await Axios.post(apiBaseURI + 'users', newUser,
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
            await Axios.delete(apiBaseURI + 'users/' + email,
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
            const res = await Axios.get<string[]>(apiBaseURI + 'credentials',
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
            await Axios.post(apiBaseURI + 'credentials', newCredentials,
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
            await Axios.delete(apiBaseURI + 'credentials/' + login,
                {
                    headers: {'Authorization': apiAccessToken}
                });
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }
}

export default EpikinsApiService;
