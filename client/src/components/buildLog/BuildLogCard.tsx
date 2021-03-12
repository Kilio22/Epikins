import React from 'react';
import { IBuildLog } from '../../interfaces/buildLog/IBuildLogInfo';
import { Card } from 'react-bootstrap';
import moment from 'moment-timezone';

interface IBuildLogCard {
    buildLog: IBuildLog
}

const BuildLogCard: React.FunctionComponent<IBuildLogCard> = ({buildLog}) => {
    const momentObj = moment.unix(buildLog.time);
    return (
        <Card className={'mt-2'}>
            <Card.Body>
                <span className={'font-weight-bold'}>{buildLog.starter.split('@')[0]}</span> started a
                build
                on <span className={'font-weight-bold'}>{buildLog.project}</span> project
                (<span className={'font-weight-bold'}>{buildLog.module}</span>) for <span
                className={'font-weight-bold'}>{buildLog.target}</span> workgroup, on{' '}
                <span
                    className={'font-weight-bold'}>{momentObj.format('ddd, MMMM Do, YYYY, HH:mm')} GMT{momentObj.format('Z')}</span>
            </Card.Body>
        </Card>
    );
};

export default BuildLogCard;
