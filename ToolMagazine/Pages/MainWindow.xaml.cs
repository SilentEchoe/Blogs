using System.Collections.Generic;
using System.Configuration;
using System.Data;
using System.IO;
using System.Windows;
using System.Xml;
using System.Xml.Linq;
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
            //List<string> list = new List<string>();
            //XmlReaderSettings settings = new XmlReaderSettings();
            //string Path = @"C:\Users\Lenovo\Desktop\toolTable.xml";
            //using (XmlReader xmlReader = XmlReader.Create(Path, settings))
            //{
            //    while (xmlReader.Read())
            //    {
            //        list.Add(xmlReader.Value);
            //    }
            //}

        }
    }
}
