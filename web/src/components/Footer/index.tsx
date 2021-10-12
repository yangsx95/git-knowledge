import { useIntl } from 'umi';
import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-layout';

export default () => {
  const intl = useIntl();
  const defaultMessage = intl.formatMessage({
    id: 'app.copyright.produced',
    defaultMessage: '杨顺翔倾情奉献，欢迎访问我的小站',
  });

  const currentYear = new Date().getFullYear();

  return (
    <DefaultFooter
      copyright={`${currentYear} ${defaultMessage}`}
      links={[
        {
          key: 'Git Knowledge',
          title: 'Git Knowledge',
          href: 'https://pro.ant.design',
          blankTarget: true,
        },
        {
          key: 'github',
          title: <GithubOutlined />,
          href: 'https://github.com/yangsx95/git-knowledge',
          blankTarget: true,
        },
        {
          key: 'Author',
          title: 'Author',
          href: 'http://yangsx95.com',
          blankTarget: true,
        },
      ]}
    />
  );
};
