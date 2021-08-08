import React from "react";
import classNames from "classnames";
import { GenericComponentProps } from "types/ComponentProps";

export default function List({
  className,
  children,
  title,
  headerComponent,
  ...restProps
}: GenericComponentProps & {
  title?: string;
  headerComponent?: any;
}) {
  return (
    <>
      {title && !headerComponent ? <List.Header>{title}</List.Header> : null}
      {!title && headerComponent ? headerComponent : null}
      <div
        className={classNames(
          "mt-4 bg-white border grid grid-cols-1 divide-y",
          className
        )}
        {...restProps}
      >
        {children}
      </div>
    </>
  );
}

List.Header = function ListHeader({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames(
        "flex flex-row items-center mt-6 mb-6 gap-4",
        className
      )}
      {...restProps}
    >
      <p className="text-xl font-semibold text-gray-400 flex-0">{children}</p>
    </div>
  );
};

List.Item = function ListItem({
  className,
  children,
  ...restProps
}: GenericComponentProps) {
  return (
    <div
      className={classNames("flex items-center justify-between p-5", className)}
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
