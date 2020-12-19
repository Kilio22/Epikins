import * as React from 'react';
import '@fortawesome/fontawesome-free/css/all.css';
import {Container, Nav, Navbar, NavDropdown} from "react-bootstrap";
import {NavLink} from "react-router-dom";
import {OnSignOutClick} from "../interfaces/Functions";
import {IUser} from "../interfaces/IUser";

interface INavbarProps {
    onSignOutClick: OnSignOutClick,
    user: IUser
}

const UserAvatar = () => {
    return (
        <i className="far fa-user-circle fa-lg"
           style={{width: '28px'}}/>
    );
}

const NavBar: React.FunctionComponent<INavbarProps> = ({user, onSignOutClick}) => {
    return (
        <div>
            <Navbar collapseOnSelect expand={"lg"} bg={"primary"} variant={"dark"}>
                <Container>
                    <Navbar.Brand as={NavLink} to={"/"}>Epikins</Navbar.Brand>
                    <Navbar.Toggle aria-controls={"responsive-navbar-nav"}/>
                    <Navbar.Collapse id={"responsive-navbar-nav"}>
                        <Nav className={"mr-auto"}>
                        </Nav>
                        <Nav>
                            <NavDropdown title={
                                <span>
                                    <UserAvatar/>{user.name}
                                </span>
                            } id={"collasible-nav-dropdown"}>
                                <NavDropdown.Header>{user.email}</NavDropdown.Header>
                                <NavDropdown.Divider/>
                                <NavDropdown.Item onClick={onSignOutClick}>
                                    <i className="fas fa-sign-out-alt"/> Sign out</NavDropdown.Item>
                            </NavDropdown>
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        </div>
    );
}

export default NavBar;