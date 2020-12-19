import React, {BaseSyntheticEvent} from "react";
import {
    IProjectJobsRendererProps,
    IProjectJobsRendererState,
    jobsRendererInitialState
} from "../../interfaces/jobs/IProjectJobsRenderer";
import Fuse from "fuse.js";
import {TextField} from "@material-ui/core";
import {IGroupData} from "../../interfaces/IGroupData";
import Legend from "./Legend";
import BuildToolbox from "./BuildToolbox";
import GroupsData from "./GroupsData";
import {Form} from "react-bootstrap";
import {appInitialContext} from "../../interfaces/IAppContext";

const jobsFuseOptions: Fuse.IFuseOptions<IGroupData> = {
    shouldSort: true,
    threshold: 0.4,
    keys: ["groupJob.job.name"]
};

class ProjectJobsRenderer extends React.Component<IProjectJobsRendererProps, IProjectJobsRendererState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>

    constructor(props: IProjectJobsRendererProps) {
        super(props);

        this.state = jobsRendererInitialState;
        this.onSearchFieldChange = this.onSearchFieldChange.bind(this);
        this.onJobClick = this.onJobClick.bind(this);
    }

    render() {
        let groupsData: IGroupData[] = this.props.groupsData;
        const fuse = new Fuse(this.props.groupsData, jobsFuseOptions);

        if (this.state.queryString !== "") {
            const fuseResult = fuse.search(this.state.queryString);

            groupsData = fuseResult.map(fuseRes => {
                return fuseRes.item;
            });
        }
        return (
            <div>
                <TextField placeholder={"Group name"} variant={"standard"}
                           color={"primary"}
                           onChange={(event => this.onSearchFieldChange(event.target.value.trim()))}
                           className={"ml-1"}
                           autoFocus={true}/>
                <BuildToolbox selectedJobs={this.props.selectedJobs} isBuilding={this.props.isBuilding}
                              onBuildClick={this.props.onBuildClick}
                              onGlobalBuildClick={this.props.onGlobalBuildClick}/>
                <Form className={"fu-switch mt-0"}>
                    <Form.Check
                        type={"switch"}
                        id={"custom-switch"}
                        label={"Follow-up"}
                        checked={this.context.fuMode}
                        onChange={() => this.context.changeAppStateByProperty &&
                            this.context.changeAppStateByProperty("fuMode", !this.context.fuMode, false)}
                    />
                </Form>
                <GroupsData groupsData={groupsData} selectedJobs={this.props.selectedJobs}
                            onCheckboxChange={this.props.onCheckboxChange} onJobClick={this.onJobClick}/>
                <Legend/>
            </div>
        );
    }

    onSearchFieldChange(value: string) {
        this.setState({
            queryString: value
        });
    }

    onJobClick(event: BaseSyntheticEvent, url: string) {
        const target = event.target as HTMLInputElement;
        if (target.type !== "checkbox") {
            window.open(url);
        }
    }
}

export default ProjectJobsRenderer