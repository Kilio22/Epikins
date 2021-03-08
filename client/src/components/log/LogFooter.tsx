import React from 'react';
import ReactPaginate from 'react-paginate';

type OnPageClick = (selectedItem: { selected: number }) => void;

interface ILogFooterProps {
    totalPage: number,
    onPageClick: OnPageClick
}

const LogFooter: React.FunctionComponent<ILogFooterProps> = ({totalPage, onPageClick}) => {
    return (
        <ReactPaginate
            previousLabel={'Previous'}
            nextLabel={'Next'}
            breakLabel={'...'}
            pageCount={totalPage}
            marginPagesDisplayed={2}
            pageRangeDisplayed={5}
            onPageChange={onPageClick}
            activeClassName={'active'}
        />
    );
};

export default LogFooter;
