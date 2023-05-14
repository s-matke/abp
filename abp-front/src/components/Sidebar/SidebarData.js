import React from "react";
import * as FaIcons from "react-icons/fa";
import * as AiIcons from "react-icons/ai";
import * as IoIcons from "react-icons/io";
import * as RiIcons from "react-icons/ri";
 
export const SidebarData = [
  {
    title: "Workers",
    role: ['Admin'],
    icon: <FaIcons.FaUserCog />,
    iconClosed: <RiIcons.RiArrowDownSFill />,
    iconOpened: <RiIcons.RiArrowUpSFill />,
 
    subNav: [
      {
        title: "Create",
        path: "/worker/create",
        icon: <IoIcons.IoMdCreate />,
        role: ['Admin']
      },
      {
        title: "Search",
        path: "/worker/search",
        icon: <IoIcons.IoIosSearch />,
        role: ['Admin']
      },
    ],
  },
  {
    title: "Production",
    role: ['Worker', 'User', 'Plan Manager'],
    icon: <FaIcons.FaUserCog />,
    iconClosed: <RiIcons.RiArrowDownSFill />,
    iconOpened: <RiIcons.RiArrowUpSFill />,
 
    subNav: [
      {
        title: "Plan",
        path: "/production/plan",
        icon: <IoIcons.IoMdCreate />,
        role: ['Worker', 'Plan Manager']
      },
      {
        title: "Search",
        path: "/production/search",
        icon: <IoIcons.IoIosSearch />,
        role: ['User', 'Worker', 'Plan Manager']
      },
    ],
  },
  {
    title: "Product",
    role: ['Worker', 'Plan Manager', 'Inventory Manager'],
    icon: <FaIcons.FaCarBattery />,
    iconClosed: <RiIcons.RiArrowDownSFill />,
    iconOpened: <RiIcons.RiArrowUpSFill />,
 
    subNav: [
      {
        title: "Add",
        path: "/product/create",
        icon: <IoIcons.IoMdCreate />,
        role: ['Worker', 'Inventory Manager']
      },
      {
        title: "Search",
        path: "/product/search",
        icon: <IoIcons.IoIosSearch />,
        role: ['User', 'Worker', 'Plan Manager', 'Inventory Manager']
      },
    ],
  },
  {
    title: "Material",
    role: ['Worker', 'Plan Manager', 'Inventory Manager'],
    icon: <FaIcons.FaTools />,
    iconClosed: <RiIcons.RiArrowDownSFill />,
    iconOpened: <RiIcons.RiArrowUpSFill />,
 
    subNav: [
      {
        title: "Add",
        path: "/material/create",
        icon: <IoIcons.IoMdCreate />,
        role: ['Worker', 'Inventory Manager']
      },
      {
        title: "Search",
        path: "/material/search",
        icon: <IoIcons.IoIosSearch />,
        role: ['User', 'Worker', 'Plan Manager', 'Inventory Manager']
      },
    ],
  },
  {
    title: "Sign Up",
    path: "/signup",
    role: ['unknown'],
    icon: <FaIcons.FaUserAlt />,
  },
  {
    title: "Sign In",
    path: "/signin",
    role: ['unknown'],
    icon: <FaIcons.FaSignInAlt />,//<IoIcons.IoIosLogIn />,
  },
  {
    title: "Sign Out",
    path: "/signout",
    cName: "nav-text",
    role: ["Admin", "User", "Worker", "Plan Manager", 'Inventory Manager']
}
];