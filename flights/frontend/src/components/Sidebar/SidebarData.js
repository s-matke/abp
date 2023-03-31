import React from "react";
import * as FaIcons from "react-icons/fa";
import * as AiIcons from "react-icons/ai";
import * as IoIcons from "react-icons/io";
import * as RiIcons from "react-icons/ri";
 
export const SidebarData = [
  {
    title: "Flights",
    role: ['guest'],
    icon: <FaIcons.FaPlaneDeparture />,
    iconClosed: <RiIcons.RiArrowDownSFill />,
    iconOpened: <RiIcons.RiArrowUpSFill />,
 
    subNav: [
      {
        title: "Create",
        path: "/flight/create",
        icon: <IoIcons.IoMdCreate />,
        role: ['admin']
      },
      {
        title: "Search",
        path: "/flight/search",
        icon: <IoIcons.IoIosSearch />,
        role: ['guest', 'user', 'admin']
      },
    ],
  },
  {
    title: "Sign in",
    path: "/signin",
    role: ['guest'],
    icon: <FaIcons.FaSignInAlt />,//<IoIcons.IoIosLogIn />,
  }
];