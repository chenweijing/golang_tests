#include <linux/kernel.h>
#include <linux/ip.h>
#include <linux/version.h>
#include <linux/netfilter.h>
#include <linux/netfilter_ipv4.h>
#include <linux/skbuff.h>
#include <linux/netfilter_ipv4/ip_tables.h>
#include <linux/moduleparam.h>
#include <linux/in.h>
#include <linux/socket.h>
#include <linux/icmp.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("ZHT");
MODULE_DESCRIPTION("My Hook Test");

static int pktcnt = 0;
static unsigned int myhook_func(unsigned int hooknum, struct sk_buff **skb,
                               const struct net_device *in,
                               const struct net_device *out,
                               int (*okfn)(struct sk_buff *)) {
	struct iphdr *ip_hdr = (struct iphdr *)skb_network_header(skb);
	printk ("%u.%u.%u.%u\n",NIPQUAD(ip_hdr->daddr));
	return NF_ACCEPT;
}

static struct nf_hook_ops nfho = {
	.hook = myhook_func,
	.owner = THIS_MODULE,
	.pf = PF_INET,
	.hooknum = 3,
	.priority = NF_IP_PRI_FIRST,
};

static int __init myhook_init(void) {
	nf_register_hook(&nfho);
}

static void __exit myhook_finit(void) {
	nf_unregister_hook(&nfho);
}

module_init(myhook_init);
module_exit(myhook_finit);