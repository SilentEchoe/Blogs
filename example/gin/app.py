import gradio as gr
import torch
from transformers import DistilBertTokenizer, DistilBertForSequenceClassification

# 加载模型和分词器
tokenizer = DistilBertTokenizer.from_pretrained('distilbert-base-uncased-finetuned-sst-2-english')
model = DistilBertForSequenceClassification.from_pretrained('distilbert-base-uncased-finetuned-sst-2-english')

# 模型函数，接受输入并返回输出
def predict(text):
    # 分词
    inputs = tokenizer.encode_plus(
        text,
        None,
        add_special_tokens=True,
        max_length=64,
        pad_to_max_length=True,
        return_token_type_ids=True,
        truncation=True
    )

    # 将分词后的输入转换为张量
    input_ids = torch.tensor(inputs['input_ids']).unsqueeze(0)
    attention_mask = torch.tensor(inputs['attention_mask']).unsqueeze(0)

    # 将输入传递给模型进行预测
    outputs = model(input_ids, attention_mask=attention_mask)
    prediction = torch.argmax(outputs[0])

    # 返回模型的预测结果
    if prediction.item() == 0:
        return "负面情绪"
    else:
        return "正面情绪"

# 创建一个接口，将输入与输出连接起来
interface = gr.Interface(fn=predict, inputs="text", outputs="text")

# 启动接口
interface.launch()