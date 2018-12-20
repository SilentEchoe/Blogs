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
                if (cn.State != System.Data.ConnectionState.Open)
                {
                    cn.Open();
                    SQLiteCommand cmd = new SQLiteCommand();
                    cmd.Connection = cn;
                    // 创建表结构
                    cmd.CommandText = "CREATE TABLE t1(id varchar(4),score int)";
                    //cmd.CommandText = "CREATE TABLE IF NOT EXISTS t1(id varchar(4),score int)";
                    cmd.ExecuteNonQuery();
                }
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
