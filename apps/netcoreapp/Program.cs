var builder = WebApplication.CreateBuilder(args);

builder.Services.AddEndpointsApiExplorer();
builder.WebHost.UseUrls("http://+:8080");
var app = builder.Build();

var list = new[]
{
    "Hello", "From", "NetCore6"
};

app.MapGet("/", () =>
{
    
    return list.ToList();
})
.WithName("Get");

app.Run();