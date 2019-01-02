using Stylet;
using StyletIoC;
using ToolMagazine.Pages;

namespace ToolMagazine
{
    public class Bootstrapper : Bootstrapper<IndexViewModel>
    {
        protected override void ConfigureIoC(IStyletIoCBuilder builder)
        {
            // Configure the IoC container in here
        }

        protected override void Configure()
        {
            // Perform any other configuration before the application starts
        }
    }
}
