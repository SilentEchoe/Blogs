<template>
    <div>
        <div id="Top">
            <Button>Default</Button>
            顶部
        </div>
        
        <div id="CenTen">中部内容</div>
        
        <div id="Bottom">底部内容</div>
       
    </div>
</template>



<script>
import { Button, Table } from 'view-design';

export default {
    
}
</script>

<style  scoped>
div{
    border-style: solid; 
    border-width:5px;
}

#Top{   
    width: 80%;
    height: 100px;
    border:red 1px;
}

#CenTen{
    width: 100%;
    height: 70%;
    
}

#Bottom{
    width: 100%;
    height: 10%;
    
}

</style>>

</style>