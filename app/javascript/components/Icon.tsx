import React from "react";
import classNames from "classnames";
import { GenericComponentProps } from "types/ComponentProps";

export default function Icon({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div className={classNames(className)} {...restProps}>
      {children}
    </div>
  );
}

Icon.Wallet = function IconWallet({
  className,
  ...restProps
}: {
  className?: string;
}) {
  return (
    <Icon className={classNames(className)} {...restProps}>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        width="24"
        height="24"
      >
        <path fill="none" d="M0 0h24v24H0z" />
        <path
          d="M18 7h3a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h15v4zM4 9v10h16V9H4zm0-4v2h12V5H4zm11 8h3v2h-3v-2z"
          fill="rgba(31,41,55,1)"
        />
      </svg>
    </Icon>
  );
};

Icon.File = function IconFile({
  className,
  ...restProps
}: {
  className?: string;
}) {
  return (
    <Icon className={classNames(className)} {...restProps}>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        width="24"
        height="24"
      >
        <path fill="none" d="M0 0h24v24H0z" />
        <path
          d="M19 22H5a3 3 0 0 1-3-3V3a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v12h4v4a3 3 0 0 1-3 3zm-1-5v2a1 1 0 0 0 2 0v-2h-2zm-2 3V4H4v15a1 1 0 0 0 1 1h11zM6 7h8v2H6V7zm0 4h8v2H6v-2zm0 4h5v2H6v-2z"
          fill="rgba(31,41,55,1)"
        />
      </svg>
    </Icon>
  );
};

Icon.Empty = function IconEmpty({
  className,
  width,
  height,
  color,
  ...restProps
}: {
  className?: string;
  width?: string;
  height?: string;
  color?: string;
}) {
  return (
    <Icon className={classNames(className)} {...restProps}>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        width={width || "24"}
        height={width || "24"}
      >
        <path fill="none" d="M0 0H24V24H0z" />
        <path
          d="M11 11v2l-5.327 6H11v2H3v-2l5.326-6H3v-2h8zm10-8v2l-5.327 6H21v2h-8v-2l5.326-6H13V3h8z"
          fill={color || "rgba(0,0,0,1)"}
        />
      </svg>
    </Icon>
  );
};
