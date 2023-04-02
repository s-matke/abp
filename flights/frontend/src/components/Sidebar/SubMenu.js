import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import styled from "styled-components";
 
const SidebarLink = styled(Link)`
  display: flex;
  color: #e1e9fc;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  list-style: none;
  height: 60px;
  text-decoration: none;
  font-size: 18px;
 
  &:hover {
    background: #252831;
    border-left: 4px solid white;
    cursor: pointer;
  }
`;
 
const SidebarLabel = styled.span`
  margin-left: 16px;
`;
 
const DropdownLink = styled(Link)`
  background: #252831;
  height: 60px;
  padding-left: 3rem;
  display: flex;
  align-items: center;
  text-decoration: none;
  color: #f5f5f5;
  font-size: 18px;
 
  &:hover {
    background: gray;
    cursor: pointer;
  }
`;
 
const SubMenu = ({ item }) => {
  const [subnav, setSubnav] = useState(false);
 
  const showSubnav = () => setSubnav(!subnav);

  const [userRole, setUserRole] = useState(localStorage.getItem('user') || 'guest')
  console.log(userRole)
  

  function updateSidebarData() {
    const role = localStorage.getItem('user') || 'guest';
    if (role !== "guest") {
      let parsedUser = JSON.parse(role)
      setUserRole(parsedUser['Role'])
    }
  }


  
  useEffect(() => {
    updateSidebarData()  
  }, [userRole])
   
  return (
    <>
      <SidebarLink to={item.path}
      onClick={item.subNav && showSubnav}>
        <div>
          {item.icon}
          <SidebarLabel>{item.title}</SidebarLabel>
        </div>
        <div>
          {item.subNav && subnav
            ? item.iconOpened
            : item.subNav
            ? item.iconClosed
            : null}
        </div>
      </SidebarLink>
      {subnav && 
        item.subNav.map((item, index) => {
            if (item.role.includes(userRole.toLowerCase())) {
                return (
                    <DropdownLink to={item.path} key={index}>
                        {item.icon}
                        <SidebarLabel>{item.title}</SidebarLabel>
                    </DropdownLink>
                );
            }  
        })}
    </>
  );
};
 
export default SubMenu;