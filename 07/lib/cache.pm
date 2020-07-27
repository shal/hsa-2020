package cache;

use warnings;

use File::Path;

sub purge {
    my $r = shift;

    $r->send_http_header("text/html");
    return OK if $r->header_only;

    opendir(DH, "/var/cache/nginx");
    my @dirs = readdir(DH);
    closedir(DH);

    foreach my $dir (@dirs)
    {
        next if($dir =~ /^\.$/);
        next if($dir =~ /^\.\.$/);
        next if($dir =~ /^client_temp$/);
        next if($dir =~ /^fastcgi_temp$/);
        next if($dir =~ /^proxy_temp$/);
        next if($dir =~ /^scgi_temp$/);
        next if($dir =~ /^uwsgi_temp$/);

        $dir_path = "/var/cache/nginx/${dir}";
        rmtree $dir_path;
        $r->print($dir_path, "\n");
    }

    if (-f $r->filename or -d _) {
        $r->print($r->uri, " exists!\n");
    }

    return OK;
}

1;
__END__