import React, {useState} from 'react';
import {Button, Card, Col, Input, Row, Select} from 'antd';
import {genApi} from "../service";
import {sync} from "../util";

const {TextArea} = Input;
const { Option } = Select;

export default () => {
  const [schema, setSchema] = useState('');
  const [struct, setStruct] = useState('');
  const [tags, setTags] = useState(["json"]);

  const onChange = ({target: {value}}) => {
    setSchema(value);
  };

  const handleTags = (value) => {
    setTags(value);
  };

  const onConvert = () => {
    if (schema === "") {
      return
    }
    sync(async function () {
      const ret = await genApi({"table": schema, "tags": tags})
      setStruct(ret);
    })
  };

  return (
    <div style={{margin: '20px'}}>
      <Card
        title="MySQL 建表语句转换成 Golang 结构体"
        extra={
          <Button type="primary" onClick={onConvert}>
            开始转换
          </Button>
        }
      >
        <Row style={{marginBottom:'24px'}}>
        <div style={{width:'100px',fontSize:'15px'}}>Tag：</div>
          <Select mode="tags" style={{width: '100%'}} defaultValue={tags} placeholder="Tags" onChange={handleTags}>
            {tags}
          </Select>
        </Row>
        <Row gutter={{xs: 8, sm: 16, md: 24, lg: 32}}>
          <Col span={12}>
            <TextArea
              onChange={onChange}
              placeholder="输入 MySQL 建表语句"
              autoSize={{minRows: 20, maxRows: 20}}
            />
          </Col>
          <Col span={12}>
            <TextArea
              value={struct}
              readOnly
              autoSize={{minRows: 20, maxRows: 20}}
            />
          </Col>
        </Row>
      </Card>
    </div>
  );
};
