#include "osinfo.h"
#ifdef _WIN32
#include <Windows.h>

char *GetOsVersion()
{
    OSVERSIONINFO osvi;
    ZeroMemory(&osvi, sizeof(OSVERSIONINFO));
    osvi.dwOSVersionInfoSize = sizeof(OSVERSIONINFO);
    char *osName = (char *)malloc(100 * sizeof(char));
    char *osVersion = (char *)malloc(100 * sizeof(char));

    if (GetVersionEx(&osvi))
    {
        // Windows版本名
        switch (osvi.dwMajorVersion)
        {
        case 10:
            if (osvi.dwMinorVersion == 0)
            {
                sprintf(osName, "Windows 10");
            }
            else
            {
                sprintf(osName, "Windows 10 or later (unknown minor version)");
            }
            break;
        case 6:
            switch (osvi.dwMinorVersion)
            {
            case 3:
                sprintf(osName, "Windows 8.1");
                break;
            case 2:
                sprintf(osName, "Windows 8");
                break;
            case 1:
                sprintf(osName, "Windows 7");
                break;
            case 0:
                sprintf(osName, "Windows Vista");
                break;
            default:
                sprintf(osName, "Unknown Windows version");
                break;
            }
            break;
        case 5:
            switch (osvi.dwMinorVersion)
            {
            case 2:
                sprintf(osName, "Windows Server 2003 or Windows XP x64 Edition");
                break;
            case 1:
                sprintf(osName, "Windows XP");
                break;
            case 0:
                sprintf(osName, "Windows 2000");
                break;
            default:
                sprintf(osName, "Unknown Windows version");
                break;
            }
            break;
        default:
            sprintf(osName, "Unknown Windows version");
            break;
        }
        // If Win11
        if (osvi.dwBuildNumber >= 22000)
        {
            sprintf(osName, "Windows 11");
        }
        sprintf(osVersion, "%s (Build %d)", osName, osvi.dwBuildNumber);
        return osVersion;
    }
    else
    {
        printf("Error getting Windows version information\n");
        return "Error getting Windows version information";
    }
}
#else
char *GetOsVersion()
{
    printf("Unknown OS Version\n");
    return "Unknown OS Version";
}
#endif