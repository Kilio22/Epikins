import Axios, {AxiosResponse} from "axios";
import {apiBaseURI} from "../Config";
import {IGroupData} from "../interfaces/IGroupData";
import {IJob} from "../interfaces/IJob";

class EpikinsApiService {
    static async login(accessToken: string): Promise<boolean> {
        try {
            await Axios.post(apiBaseURI + "login", {}, {
                headers: {
                    "Authorization": accessToken,
                }
            });
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }

    static async getGroupsData(url: string, apiAccessToken: string): Promise<IGroupData[] | null> {
        try {
            const res: AxiosResponse<IGroupData[]> = await Axios.get<IGroupData[]>(url, {headers: {"Authorization": apiAccessToken}});
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async getJobs(url: string, apiAccessToken: string): Promise<IJob[] | null> {
        try {
            const res: AxiosResponse<IJob[]> = await Axios.get<IJob[]>(url, {headers: {"Authorization": apiAccessToken}});
            return res.data;
        } catch (e) {
            console.log(e);
            return null;
        }
    }

    static async buildJobs(requestedBuilds: string[], project: string, visibility: string, fuMode: boolean, apiAccessToken: string): Promise<boolean> {
        try {
            await Axios.post(apiBaseURI + "build",
                requestedBuilds,
                {
                    params: {"visibility": visibility, "project": project, "fu": fuMode},
                    headers: {"Authorization": apiAccessToken}
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
            await Axios.post(apiBaseURI + "build/global", {},
                {
                    params: {"visibility": visibility, "project": project},
                    headers: {"Authorization": apiAccessToken}
                }
            );
            return true;
        } catch (e) {
            console.log(e);
        }
        return false;
    }
}

export default EpikinsApiService;