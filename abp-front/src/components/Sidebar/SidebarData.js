import React from "react";
import * as FaIcons from "react-icons/fa";
import * as AiIcons from "react-icons/ai";
import * as IoIcons from "react-icons/io";
import * as RiIcons from "react-icons/ri";
 
export const SidebarData = [
  {
    title: "Accommodation",
    role: ['host', 'admin', 'guest', 'unknown'],
    icon: <FaIcons.FaBed />,
    iconClosed: <RiIcons.RiArrowDownSFill />,
    iconOpened: <RiIcons.RiArrowUpSFill />,
 
    subNav: [
      {
        title: "Add",
        path: "/accommodation/create",
        icon: <IoIcons.IoMdCreate />,
        role: ['host', 'admin']
      },
      {
        title: "Search",
        path: "/accommodation/search",
        icon: <IoIcons.IoIosSearch />,
        role: ['guest', 'host', 'admin', 'unknown']
      },
      {
        title: "Set Price",
        path: "/accommodation",
        icon: <IoIcons.IoIosSearch />,
        role: ['host']
      },
      {
        title: "Pending Reservations",
        path: "/reservation",
        icon: <IoIcons.IoIosSearch />,
        role: ['host']
      },
    ],
  },
  {
    title: "User info",
    path: "/userupdate",
    role: ['host', 'guest'],
    icon: <FaIcons.FaUserAlt />,
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
    role: ["host", "guest"]
}
];