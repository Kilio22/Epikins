import React from 'react';
import ReactPaginate from 'react-paginate';

type OnPageClick = (selectedItem: { selected: number }) => void;

interface ILogFooterProps {
    currentPage: number,
    totalPage: number,
    onPageClick: OnPageClick
}

const LogFooter: React.FunctionComponent<ILogFooterProps> = ({currentPage, totalPage, onPageClick}) => {
    return (
        <ReactPaginate
            forcePage={currentPage - 1}
            previousLabel='&larr;'
            nextLabel='&rarr;'
            containerClassName={'pagination d-flex justify-content-center mt-4'}
            pageClassName={'page-item'}
            pageLinkClassName={'page-link'}
            activeClassName={'page-item active'}
            activeLinkClassName={'page-link'}
            disabledClassName={'page-item disabled'}
            nextClassName={'page-item'}
            nextLinkClassName={'page-link'}
            previousClassName={'page-item'}
            previousLinkClassName={'page-link'}
            pageCount={totalPage}
            marginPagesDisplayed={2}
            pageRangeDisplayed={5}
            onPageChange={onPageClick}
        />
    );
};

export default LogFooter;
