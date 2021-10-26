import {
  BookOutlined,
  LockOutlined,
  MailOutlined,
  UserOutlined,
} from '@ant-design/icons';
import {Alert, message} from 'antd';
import React, {useState} from 'react';
import ProForm, {ProFormText} from '@ant-design/pro-form';
import {useIntl, history, FormattedMessage, SelectLang, Link} from 'umi';
import Footer from '@/components/Footer';

import styles from './index.less';
import {register} from "@/services/login";
import type {API} from "@/services/login/typing";

const LoginMessage: React.FC<{
  content: string;
}> = ({content}) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

const Login: React.FC = () => {
  // 表单的提交状态
  const [submitting, setSubmitting] = useState(false);
  const [userRegisterState, setUserRegisterState] = useState<API.RegisterResult>({
    code: 0,
    data: undefined,
    detail: "",
    msg: ""
  });

  const intl = useIntl();

  const handleSubmit = async (values: API.RegisterParams) => {
    setSubmitting(true); // 设置表单状态为提交中
    try {
      // 发送注册接口
      const resp = await register(values);

      if (resp.code === 200) {
        const defaultLoginSuccessMessage = intl.formatMessage({
          id: 'pages.register.success',
          defaultMessage: '注册成功！',
        });
        message.success(defaultLoginSuccessMessage);
        /** 此方法会跳转到 redirect 参数所在的位置 */
        if (!history) return;
        const {query} = history.location;
        const {redirect} = query as { redirect: string };
        history.push(redirect || '/');
        return;
      } else if (resp.code == 440) { // 用户id已存在
        message.error(intl.formatMessage({
          id: 'pages.register.userid.exists',
          defaultMessage: '用户ID已经存在',
        }));
      } else if (resp.code == 441) { // 用户邮箱已被注册
        message.error(intl.formatMessage({
          id: 'pages.register.email.exists',
          defaultMessage: '邮箱地址已经存在',
        }));
      }
      // 如果失败去设置用户错误信息
      setUserRegisterState({...resp});
    } catch (error) {
      console.log(error)
      const defaultLoginFailureMessage = intl.formatMessage({
        id: 'pages.register.failure',
        defaultMessage: '注册失败，请重试！',
      });
      message.error(defaultLoginFailureMessage);
    }
    setSubmitting(false);
  };
  const {code, detail} = userRegisterState;

  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang/>}
      </div>
      <div className={styles.content}>
        <div className={styles.top}>
          <div className={styles.header}>
            <Link to="/">
              <span className={styles.title}>创建一个GitKnowledge ID</span>
            </Link>
          </div>
        </div>
        <div className={styles.main}>
          <ProForm
            initialValues={{
              autoLogin: true,
            }}
            submitter={{
              searchConfig: {
                submitText: intl.formatMessage({
                  id: 'pages.register.submit',
                  defaultMessage: '注册',
                }),
              },
              render: (_, dom) => dom.pop(),
              submitButtonProps: {
                loading: submitting,
                size: 'large',
                style: {
                  width: '100%',
                },
              },
            }}
            onFinish={async (values) => {
              await handleSubmit(values as API.RegisterParams);
            }}
          >

            {code !== 200 && (
              <LoginMessage
                content={intl.formatMessage({
                  id: 'pages.register.errorMessage',
                  defaultMessage: '注册失败',
                })}
              />
            ) && console.log(detail)}

            <>
              <ProFormText
                name="userid"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={styles.prefixIcon}/>,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.register.userid.placeholder',
                  defaultMessage: 'GitKnowledge ID',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.userid.validate"
                        defaultMessage="用户ID不符合规则，必须是数字和字母!"
                      />
                    ),
                    pattern: /^[a-zA-Z0-9_]{4,16}$/
                  },
                ]}
              />
              <ProFormText
                name="nickname"
                fieldProps={{
                  size: 'large',
                  prefix: <BookOutlined className={styles.prefixIcon}/>,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.register.username.placeholder',
                  defaultMessage: '请输入用户名',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="用户名不符合规则!"
                      />
                    ),
                    pattern: /^[a-zA-Z0-9_]{4,16}$/
                  },
                ]}
              />
              <ProFormText.Password
                name="password"
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={styles.prefixIcon}/>,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.register.password.placeholder',
                  defaultMessage: '密码: ant.design',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.login.password.required"
                        defaultMessage="请输入密码！"
                      />
                    ),
                    pattern: /^.*(?=.{6,})(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*? ]).*$/
                  },
                ]}
              />
            </>
            <ProFormText
              name="email"
              fieldProps={{
                size: 'large',
                prefix: <MailOutlined className={styles.prefixIcon}/>,
              }}
              placeholder={intl.formatMessage({
                id: 'pages.register.email.placeholder',
                defaultMessage: '请输入您的邮箱地址',
              })}
              rules={[
                {
                  required: true,
                  message: (
                    <FormattedMessage
                      id="pages.register.email.required"
                      defaultMessage="请输入正确的邮箱"
                    />
                  ),
                  pattern: /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$/
                },
              ]}
            />
          </ProForm>
        </div>
      </div>
      <Footer/>
    </div>
  );
};

export default Login;
