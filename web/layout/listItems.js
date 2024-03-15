import * as React from 'react';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import DashboardIcon from '@mui/icons-material/Dashboard';
import PowerIcon from '@mui/icons-material/Power';
import SavingsIcon from '@mui/icons-material/Savings';
import PeopleIcon from '@mui/icons-material/People';
import BarChartIcon from '@mui/icons-material/BarChart';
import LayersIcon from '@mui/icons-material/Layers';
import AssignmentIcon from '@mui/icons-material/Assignment';
import YouTubeIcon from '@mui/icons-material/YouTube';

import {NavLink} from "react-router-dom";

export const mainListItems = (

    <>
        <NavLink to="/"
                 className={({isActive, isPending}) => isPending ? "pending" : isActive ? "active" : ""}
                 style={{color: 'inherit', textDecoration: 'inherit'}}>
            <ListItemButton>
                <ListItemIcon>
                    <DashboardIcon/>
                </ListItemIcon>
                <ListItemText primary="Dashboard"/>
            </ListItemButton>
        </NavLink>

        <NavLink to="/ui/tasks"
                 className={({isActive, isPending}) => isPending ? "pending" : isActive ? "active" : ""}
                 style={{color: 'inherit', textDecoration: 'inherit'}}>
            <ListItemButton>
                <ListItemIcon>
                    <SavingsIcon/>
                </ListItemIcon>
                <ListItemText primary="Tasks"/>
            </ListItemButton>
        </NavLink>



        <NavLink to="/ui/reports"
                 className={({isActive, isPending}) => isPending ? "pending" : isActive ? "active" : ""}
                 style={{color: 'inherit', textDecoration: 'inherit'}}>
            <ListItemButton>
                <ListItemIcon>
                    <BarChartIcon/>
                </ListItemIcon>
                <ListItemText primary="Reports"/>
            </ListItemButton>
        </NavLink>

        <ListItemButton>
            <ListItemIcon>
                <LayersIcon/>
            </ListItemIcon>
            <ListItemText primary="Integrations"/>
        </ListItemButton>
    </>
);

export const secondaryListItems = (
    <>
        <ListSubheader component="div" inset>
            Saved reports
        </ListSubheader>
        <ListItemButton>
            <ListItemIcon>
                <AssignmentIcon/>
            </ListItemIcon>
            <ListItemText primary="Current month"/>
        </ListItemButton>
        <ListItemButton>
            <ListItemIcon>
                <AssignmentIcon/>
            </ListItemIcon>
            <ListItemText primary="Last quarter"/>
        </ListItemButton>
        <ListItemButton>
            <ListItemIcon>
                <AssignmentIcon/>
            </ListItemIcon>
            <ListItemText primary="Year-end sale"/>
        </ListItemButton>
    </>
);
