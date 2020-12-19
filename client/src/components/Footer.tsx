import {Container} from "react-bootstrap";
import React from "react";

const Footer = () => {
    return (
        <footer className={"footer"}>
            <Container>
                <div className={"text-right"}>
                    made with <i className="fas fa-heart heart"/> by <a
                    href={"https://github.com/Kilio22"} target={"_blank"} rel={"noopener noreferrer"}>Kylian Balan</a>
                </div>
            </Container>
        </footer>
    )
};

export default Footer;