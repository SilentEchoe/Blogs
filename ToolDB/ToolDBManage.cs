using System;
using System.Collections.Generic;
using System.Configuration;
using System.Data.SQLite;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ToolDB
{
    public class ToolDBManage
    {
        public bool CreaDB(string path)
        {
            try
            {
                path = path + ConfigurationManager.AppSettings["tooldb"];
                SQLiteConnection cn = new SQLiteConnection("data source=" + path);
                cn.Open();
                cn.Close();
                return true;
            }
            catch (Exception)
            {
                return false;

            }

        }
    }
}
