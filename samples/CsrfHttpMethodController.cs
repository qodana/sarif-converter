using Microsoft.AspNetCore.Mvc;

namespace WebApplication1.Controllers;

public class CsrfHttpMethodController : Controller
{
    [HttpGet]
    public void Get() {}

    [HttpDelete]
    public void Delete() {}

    [HttpHead]
    public void Head() {}

    [HttpOptions]
    public void Options() {}

    [HttpPatch]
    public void Patch() {}

    [HttpPost]
    public void Post() {}

    [HttpPut]
    public void Put() {}

    public IActionResult Index() => new ContentResult {Content = "Hello"};
}