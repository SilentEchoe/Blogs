using System.Configuration;
using System.IO;
using System.Windows;
using ToolDB;

namespace ToolMagazine.Pages
{
    /// <summary>
    /// MainWindow.xaml 的交互逻辑
    /// </summary>
    public partial class ShellView : Window
    {
        public ShellView()
        {
            InitializeComponent();
        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {
            MessageBox.Show("a");
        }


 

        private void Window_Loaded(object sender, RoutedEventArgs e)
        {
            ToolDBManage toolDB = new ToolDBManage();

            // string xmlPath = ConfigurationManager.AppSettings["tooldb"];
            // sqllite 数据库的位置

            string path = Directory.GetCurrentDirectory();
            string toolpath = path + ConfigurationManager.AppSettings["tooldb"];
         

            //如果数据库不存在 则创建
            if (!File.Exists(toolpath))
            {
                toolDB.CreaDB(path);
            }



            //MessageBox.Show(path);
        }
    }
}
