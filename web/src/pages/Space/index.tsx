import React from 'react';
import {Card, Alert, Row, Col} from 'antd';
import {useIntl} from 'umi';
import SpaceCard from "@/pages/Space/components/SpaceCard";

export default (): React.ReactNode => {
  const intl = useIntl();
  return (
    <>
      <Card>
        <Alert
          message={intl.formatMessage({
            id: 'pages.welcome.alertMessage',
            defaultMessage: 'Faster and stronger heavy-duty components have been released.',
          })}
          type="success"
          showIcon
          banner
          style={{
            margin: -12,
            marginBottom: 48,
          }}
        />

        <div className="site-card-wrapper">
          <Row gutter={[10, 10]}>
            <Col xs={24} sm={12} md={12} lg={8} xl={6}>
              <SpaceCard title={"Card title"} description={"This is the description"}
                         ownerAvatarUrl={"https://joeschmoe.io/api/v1/random"}
                         loading={false}
                         coverUrl={"https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"}/>
            </Col>
            <Col xs={24} sm={12} md={12} lg={8} xl={6}>
              <SpaceCard title={"Card title"} description={"This is the description"}
                         ownerAvatarUrl={"https://joeschmoe.io/api/v1/random"} loading={false}
                         coverUrl={"https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"}/>
            </Col>
            <Col xs={24} sm={12} md={12} lg={8} xl={6}>
              <SpaceCard title={"Card title"} description={"This is the description"}
                         ownerAvatarUrl={"https://joeschmoe.io/api/v1/random"} loading={false}
                         coverUrl={"https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"}/>
            </Col>
            <Col xs={24} sm={12} md={12} lg={8} xl={6}>
              <SpaceCard title={"Card title"} description={"This is the description"}
                         ownerAvatarUrl={"https://joeschmoe.io/api/v1/random"} loading={false}
                         coverUrl={"https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"}/>
            </Col>
            <Col xs={24} sm={12} md={12} lg={8} xl={6}>
              <SpaceCard title={"Card title"} description={"This is the description"}
                         ownerAvatarUrl={"https://joeschmoe.io/api/v1/random"} loading={false}
                         coverUrl={"https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"}/>
            </Col>
          </Row>
        </div>

      </Card>
    </>
  );
};
