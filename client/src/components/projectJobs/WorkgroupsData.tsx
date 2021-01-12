import React from 'react';
import { IWorkgroupsData } from '../../interfaces/IWorkgroupsData';
import { OnCheckboxChange, OnJobClick } from '../../interfaces/Functions';
import { Button, Col, Row } from 'react-bootstrap';
import { appInitialContext } from '../../interfaces/IAppContext';

interface IStudentRemainingBuildsProps {
    groupData: IWorkgroupsData
}

interface IGroupDataProps {
    workgroupData: IWorkgroupsData,
    selectedJobs: string[],
    onCheckboxChange: OnCheckboxChange,
    onJobClick: OnJobClick
}

interface IGroupsDataProps {
    workgroupsData: IWorkgroupsData[],
    selectedJobs: string[],
    onCheckboxChange: OnCheckboxChange,
    onJobClick: OnJobClick
}

const cssColorArray: string[] = [
    'red', 'tomato', 'orange', 'green', 'green', 'green'
];

const GroupMasterRemainingBuilds: React.FunctionComponent<IStudentRemainingBuildsProps> = ({groupData}) => {
    return (
        <appInitialContext.Consumer>
            {context => (
                <div className={'ml-auto mr-1'}>
                    [
                    {
                        context.fuMode ?
                            <span className={'font-weight-bold remaining-builds-green'}>
                                ∞
                            </span>
                            :
                            <span className={'font-weight-bold remaining-builds-' +
                            cssColorArray[groupData.mongoWorkgroupData.remainingBuilds % cssColorArray.length]}>
                                {groupData.mongoWorkgroupData.remainingBuilds}
                            </span>
                    }
                    ]
                </div>
            )}
        </appInitialContext.Consumer>
    );
};

const isCheckboxDisabled = (groupData: IWorkgroupsData, fuMode: boolean) => {
    if (fuMode) {
        return false;
    }
    return !groupData.mongoWorkgroupData.remainingBuilds;
};

const WorkgroupData: React.FunctionComponent<IGroupDataProps> = ({
                                                                     workgroupData,
                                                                     selectedJobs,
                                                                     onCheckboxChange,
                                                                     onJobClick
                                                                 }) => {
    return (
        <appInitialContext.Consumer>
            {context => (
                <Button variant={'outline-primary'}
                        className={'m-1 text-left d-flex align-items-center'}
                        block={true}
                        onClick={(event => onJobClick(event, workgroupData.mongoWorkgroupData.url))}>
                    <input
                        className={'jobs-checkbox mr-2'}
                        type={'checkbox'}
                        checked={selectedJobs.includes(workgroupData.mongoWorkgroupData.name)}
                        onChange={(event => onCheckboxChange(event.target.checked, workgroupData))}
                        disabled={isCheckboxDisabled(workgroupData, context.fuMode)}/>
                    {' '}
                    <i className={'fas fa-folder mr-1'}/> {workgroupData.mongoWorkgroupData.name}
                    <GroupMasterRemainingBuilds groupData={workgroupData}/>
                    {
                        (workgroupData.jobInfos.inQueue || workgroupData.jobInfos.lastBuild.buildInfos.building) &&
                        <i className="fas fa-clock fa-color-clock jobs-clock-icon"/>
                    }
                </Button>
            )}
        </appInitialContext.Consumer>
    );
};

const WorkgroupsData: React.FunctionComponent<IGroupsDataProps> = ({
                                                                       workgroupsData,
                                                                       selectedJobs,
                                                                       onCheckboxChange,
                                                                       onJobClick
                                                                   }) => {
    return (
        <div className={'mt-3'}>
            {
                <Row>
                    {
                        workgroupsData.map((workgroupData, id) => {
                            return (
                                <Col md={4} key={id}>
                                    <WorkgroupData workgroupData={workgroupData} selectedJobs={selectedJobs}
                                                   onCheckboxChange={onCheckboxChange}
                                                   onJobClick={onJobClick}/>
                                </Col>
                            );
                        })
                    }
                </Row>
            }
        </div>
    );
};

export default WorkgroupsData;
