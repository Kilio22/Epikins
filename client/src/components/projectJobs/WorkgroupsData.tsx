import React from 'react';
import { IWorkgroupsData } from '../../interfaces/IWorkgroupsData';
import { OnCheckboxChange, OnJobClick } from '../../interfaces/Functions';
import { Button, Col, Row } from 'react-bootstrap';

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
    'red', 'tomato', 'orange', 'green'
];

const GroupMasterRemainingBuilds: React.FunctionComponent<IStudentRemainingBuildsProps> = ({groupData}) => {
    return (
        <div className={'ml-auto mr-1'}>
            [
            {
                groupData.mongoWorkgroupData.remainingBuilds >= 4 ?
                    <span className={'font-weight-bold remaining-builds-green'}>
                                {groupData.mongoWorkgroupData.remainingBuilds}
                            </span>
                    :
                    <span className={'font-weight-bold remaining-builds-' +
                    cssColorArray[groupData.mongoWorkgroupData.remainingBuilds]}>
                                {groupData.mongoWorkgroupData.remainingBuilds}
                            </span>
            }
            ]
        </div>
    );
};

const WorkgroupData: React.FunctionComponent<IGroupDataProps> = ({
                                                                     workgroupData,
                                                                     selectedJobs,
                                                                     onCheckboxChange,
                                                                     onJobClick
                                                                 }) => {
    return (
        <Button variant={'outline-primary'}
                className={'m-1 text-left d-flex align-items-center'}
                block={true}
                onClick={(event => onJobClick(event, workgroupData.mongoWorkgroupData.url))}>
            <input
                className={'jobs-checkbox mr-2'}
                type={'checkbox'}
                checked={selectedJobs.includes(workgroupData.mongoWorkgroupData.name)}
                onChange={(event => onCheckboxChange(event.target.checked, workgroupData))}
                disabled={!workgroupData.mongoWorkgroupData.remainingBuilds}/>
            {' '}
            <i className={'fas fa-user-friends mr-1'}/> {workgroupData.mongoWorkgroupData.name}
            <GroupMasterRemainingBuilds groupData={workgroupData}/>
            {
                (workgroupData.jobInfos.inQueue || workgroupData.jobInfos.lastBuild.buildInfos.building) &&
                <i className="fas fa-clock fa-color-clock jobs-clock-icon"/>
            }
        </Button>
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
