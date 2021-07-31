import React from "react";
import classNames from "classnames";
import { GenericComponentProps } from "types/ComponentProps";

export default function List({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames(
        "mt-4 bg-white border grid grid-cols-1 divide-y",
        classNames
      )}
      {...restProps}
    >
      {children}
    </div>
  );
}

List.Item = function ListItem({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames(
        "flex items-center justify-between p-5",
        classNames
      )}
      {...restProps}
    >
      {children}
    </div>
  );
};

List.Icon = function ListIcon({ width, height, children, ...restProps }) {
  return (
    <div
      className="flex items-center justify-center bg-blue-400 rounded-full"
      {...restProps}
      style={{
        width,
        height,
      }}
    >
      <p className="font-semibold text-white">{children}</p>
    </div>
  );
};
