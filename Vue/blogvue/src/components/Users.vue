<template>
    <div class="layout">
        <Layout>
            <Header>
                <Menu mode="horizontal" theme="dark" active-name="1" >
                    <div class="layout-logo" >
                        <img src="logo.png" width="100" height="20" alt="logo"></div>
                    <div class="layout-nav">
                        <MenuItem name="1">
                            <Icon type="ios-navigate"></Icon>
                            Item 1
                        </MenuItem>
                        <MenuItem name="2">
                            <Icon type="ios-keypad"></Icon>
                            Item 2
                        </MenuItem>
                        <MenuItem name="3">
                            <Icon type="ios-analytics"></Icon>
                            Item 3
                        </MenuItem>
                        <MenuItem name="4">
                            <Icon type="ios-paper"></Icon>
                            Item 4
                        </MenuItem>
                    </div>
                </Menu>
            </Header>


            <Content :style="{padding: '0 50px'}">
                
                <Card style="margin-top:50px">
                    <div style="min-height:100px;">
                        Content
                    </div>
                </Card>
            </Content>

            
            <Footer class="layout-footer-center">2020-2020 &copy; TalkingData</Footer>
        </Layout>
    </div>
</template>
<script>

import { Button, Table } from 'view-design';
import markdown from '@/assets/ApiDocument.md'
// 代码高亮
import 'highlight.js/styles/github.css'
// 其他元素使用 github 的样式
import 'github-markdown-css'
export default {
    components: {
    markdown
  },
}
</script>

<style  scoped>
div{
    border-style: solid; 
    border-width:0px;
}
.layout{
    border: 1px solid #d7dde4;
    background: #f5f7f9;
    position: relative;
    border-radius: 4px;
    overflow: hidden;
    
}
.layout-logo{
    width: 100px;
    height: 30px;
    border-radius: 3px;
    float: left;
    position: relative;
    top: 15px;
    left: 20px;
    
}
.layout-nav{
    width: 420px;
    margin: 0 auto;
    margin-right: 20px;
}
.layout-footer-center{
    text-align: center;
    
}

header {
    height: 400px;
}
</style>>

</style>