import React, { BaseSyntheticEvent } from 'react';
import {
    IProjectJobsRendererProps,
    IProjectJobsRendererState,
    jobsRendererInitialState
} from '../../interfaces/jobs/IProjectJobsRenderer';
import Fuse from 'fuse.js';
import { TextField } from '@material-ui/core';
import { IWorkgroupsData } from '../../interfaces/IWorkgroupsData';
import Legend from './Legend';
import BuildToolbox from './BuildToolbox';
import { appInitialContext } from '../../interfaces/IAppContext';
import WorkgroupsData from './WorkgroupsData';

const jobsFuseOptions: Fuse.IFuseOptions<IWorkgroupsData> = {
    shouldSort: true,
    threshold: 0.4,
    keys: [ 'mongoWorkgroupData.name' ]
};

class ProjectJobsRenderer extends React.Component<IProjectJobsRendererProps, IProjectJobsRendererState> {
    static contextType = appInitialContext;
    context!: React.ContextType<typeof appInitialContext>;

    constructor(props: IProjectJobsRendererProps) {
        super(props);

        this.state = jobsRendererInitialState;
        this.onSearchFieldChange = this.onSearchFieldChange.bind(this);
        this.onJobClick = this.onJobClick.bind(this);
    }

    render() {
        const workgroupsData: IWorkgroupsData[] = this.fuseProjects([ ...this.props.workgroupsData ], this.props.workgroupsData, this.state.queryString);

        return (
            <div>
                <TextField placeholder={'Group name'} variant={'standard'}
                           color={'primary'}
                           onChange={(event => this.onSearchFieldChange(event.target.value.trim()))}
                           className={'ml-1'}
                           autoFocus={true}/>
                {
                    this.props.availableCities &&
                    <BuildToolbox availableCities={this.props.availableCities}
                                  onCitySelected={this.props.onCitySelected}
                                  selectedCity={this.props.selectedCity}
                                  selectedJobs={this.props.selectedJobs}
                                  isBuilding={this.props.isBuilding}
                                  onBuildClick={this.props.onBuildClick}
                                  onGlobalBuildClick={this.props.onGlobalBuildClick}/>
                }
                {
                    this.props.availableCities && workgroupsData.length !== 0 ?
                        <WorkgroupsData workgroupsData={workgroupsData} selectedJobs={this.props.selectedJobs}
                                        onCheckboxChange={this.props.onCheckboxChange}
                                        onJobClick={this.onJobClick}/>
                        :
                        <h2 className={'text-center'}>No jobs to display</h2>
                }
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
        if (target.type !== 'checkbox') {
            window.open(url);
        }
    }

    fuseProjects(filteredWorkgroups: IWorkgroupsData[], originalWorkgroupList: IWorkgroupsData[], queryString: string) {
        if (queryString !== '') {
            const fuse = new Fuse(originalWorkgroupList, jobsFuseOptions);
            const fuseResult = fuse.search(queryString);

            filteredWorkgroups = fuseResult.map(fuseRes => {
                return fuseRes.item;
            });
        }
        return filteredWorkgroups;
    }
}

export default ProjectJobsRenderer;
