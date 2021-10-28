import {Avatar, Card} from 'antd';
import React from 'react';
import Meta from "antd/es/card/Meta";
import {DeleteOutlined, EditOutlined, SettingOutlined} from "@ant-design/icons";

export type SpaceCardProps = {
  title: string;
  description: string;
  ownerAvatarUrl: string;
  coverUrl: string;
  loading: boolean;

};

const SpaceCard: React.FC<SpaceCardProps> = (props) => (
  <Card style={{width: 300}}
        loading={props.loading}
        cover={
          <img
            alt={props.title}
            src={props.coverUrl}
          />
        }
        actions={[
          <EditOutlined key="edit"/>,
          <SettingOutlined key="setting"/>,
          <DeleteOutlined key="delete"/>,
        ]}>
    <Meta
      avatar={<Avatar src={props.ownerAvatarUrl}/>}
      title={props.title}
      description={props.description}
    />
  </Card>
);

export default SpaceCard;
